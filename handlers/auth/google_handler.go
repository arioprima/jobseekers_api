package handlers

import (
	"encoding/json"
	"github.com/arioprima/jobseekers_api/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
)

var oauthConfig *oauth2.Config
var oauthStateString = "random"

func InitializeOAuthConfig(cfg config.Config) {
	oauthConfig = &oauth2.Config{
		ClientID:     cfg.OAuth.GoogleClientID,
		ClientSecret: cfg.OAuth.GoogleClientSecret,
		RedirectURL:  "http://localhost:8080/job-vacancies-api/auth/google/callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

func GoogleLogin(c *gin.Context) {
	url := oauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleCallback(c *gin.Context) {
	state := c.DefaultQuery("state", "")
	if state != oauthStateString {
		c.JSON(http.StatusBadRequest, gin.H{"error": "state mismatch"})
		return
	}

	code := c.DefaultQuery("code", "")
	token, err := oauthConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not get token"})
		return
	}

	client := oauthConfig.Client(c, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not get user info"})
		return
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not parse user info"})
		return
	}

	// Process user info (e.g., create or update user in your database)
	c.JSON(http.StatusOK, gin.H{"user": userInfo})
}
