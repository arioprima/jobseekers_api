package handlers

import (
	"encoding/json"
	"github.com/arioprima/jobseekers_api/config"
	"github.com/arioprima/jobseekers_api/helpers"
	"github.com/arioprima/jobseekers_api/models"
	"github.com/arioprima/jobseekers_api/schemas"
	services "github.com/arioprima/jobseekers_api/services/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"log"
	"net/http"
	"time"
)

type HandlerGoogle struct {
	Service services.ServiceGoogle
}

func NewServiceGoogleImpl(service services.ServiceGoogle) *HandlerGoogle {
	return &HandlerGoogle{Service: service}
}

var oauthConfig *oauth2.Config
var oauthStateString = "random"

func InitializeOAuthConfig(cfg config.Config) {
	oauthConfig = &oauth2.Config{
		ClientID:     cfg.OAuth.GoogleClientID,
		ClientSecret: cfg.OAuth.GoogleClientSecret,
		RedirectURL:  "http://localhost:8080/job-vacancies-api/auth/google/callback",
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

func GoogleLogin(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("id")
	if userID != nil {
		userFirstname := session.Get("firstname")
		userLastname := session.Get("lastname")
		userEmail := session.Get("email")
		userToken := session.Get("token")
		expiredAtStr := session.Get("expired_at").(string)
		var userProfileImage *string
		if img, ok := session.Get("ProfileImage").(string); ok {
			userProfileImage = &img
		}
		expiredAt, err := time.Parse(time.RFC3339, expiredAtStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse expired at"})
		}
		resData := schemas.LoginUserResponse{
			ID:           userID.(string),
			Firstname:    userFirstname.(string),
			Lastname:     userLastname.(string),
			Email:        userEmail.(string),
			RoleId:       "01908d0f-289d-7fd7-9143-d9525f8bc74d",
			RoleName:     "user",
			ProfileImage: userProfileImage,
		}

		authRes := models.TokenAuth{
			AccessToken: userToken.(string),
			Type:        "Bearer",
			ExpiredAt:   expiredAt,
		}
		helpers.ApiResponse(c, http.StatusOK, "success", "Login successfully", resData, authRes)
		return

	}
	url := oauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *HandlerGoogle) GoogleCallback(ctx *gin.Context) {
	state := ctx.DefaultQuery("state", "")
	if state != oauthStateString {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "state mismatch"})
		return
	}

	code := ctx.DefaultQuery("code", "")
	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not get token"})
		return
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not get ID token"})
		return
	}

	client := oauthConfig.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not get user info"})
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not close response body"})
		}
	}(resp.Body)

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not parse user info"})
		return
	}

	schemaDataUser := &schemas.SchemaDataUser{
		ID:           userInfo["sub"].(string),
		Firstname:    userInfo["given_name"].(string),
		Lastname:     userInfo["family_name"].(string),
		Email:        userInfo["email"].(string),
		RoleId:       "019047ca-f542-7182-8b6b-7978f905dfe7",
		ProfileImage: userInfo["picture"].(string),
	}

	authToken := models.TokenAuth{
		AccessToken: idToken,
		Type:        "Bearer",
		ExpiredAt:   time.Now().Add(time.Hour * 24),
	}

	session := sessions.Default(ctx)
	session.Set("id", schemaDataUser.ID)
	session.Set("firstname", schemaDataUser.Firstname)
	session.Set("lastname", schemaDataUser.Lastname)
	session.Set("email", schemaDataUser.Email)
	session.Set("ProfileImage", schemaDataUser.RoleId)
	session.Set("token", idToken)
	session.Set("expired_at", time.Now().Add(time.Hour*24).Format(time.RFC3339))
	err = session.Save()
	if err != nil {
		log.Println("Error saving session: ", err)
		return
	} else {
		log.Println("Session saved")
	}

	res, err := h.Service.LoginGoogleService(ctx, nil, schemaDataUser)
	if err != nil && res == nil {
		//log.Printf("Kenapa masuk ke sini: %#v", err)
		helpers.ApiResponse(ctx, http.StatusInternalServerError, "error", "Internal Server Error", nil, nil)
		return
	}
	helpers.ApiResponse(ctx, http.StatusOK, "success", "Login successfully", res, authToken)
}
