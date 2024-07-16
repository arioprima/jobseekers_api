package handlers

import (
	"github.com/arioprima/jobseekers_api/config"
	"github.com/arioprima/jobseekers_api/helpers"
	"github.com/arioprima/jobseekers_api/models"
	"github.com/arioprima/jobseekers_api/pkg"
	"github.com/arioprima/jobseekers_api/schemas"
	services "github.com/arioprima/jobseekers_api/services/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type HandlerLogin struct {
	Service services.ServiceLogin
}

func NewHandlerLogin(service services.ServiceLogin) *HandlerLogin {
	return &HandlerLogin{Service: service}
}

func (h *HandlerLogin) LoginHandler(ctx *gin.Context) {
	configs, _ := config.LoadConfig(".")
	var loginRequest schemas.SchemaDataUser

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, "error", err.Error())
		return
	}
	config := []schemas.ErrorMetaConfig{
		{
			Tag:     "required",
			Field:   "Email",
			Message: "Email is required",
		},
		{
			Tag:     "email",
			Field:   "Email",
			Message: "Email is not valid",
		},
		{
			Tag:     "required",
			Field:   "Password",
			Message: "Password is required",
		},
		{
			Tag:     "min",
			Field:   "Password",
			Message: "Password must be at least 3 characters",
			Value:   "3",
		},
	}

	errResponse, errCount := pkg.ValidatorLogin(&loginRequest, config)
	if errCount > 0 {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, "error", errResponse)
		return
	}

	res, err := h.Service.LoginService(ctx, nil, &loginRequest)

	if res.ProfileImage != nil && *res.ProfileImage == "" {
		res.ProfileImage = nil
	}

	accessTokenData := map[string]interface{}{
		"id":            res.ID,
		"email":         res.Biodata.Email,
		"firstname":     res.Biodata.Firstname,
		"lastname":      res.Biodata.Lastname,
		"role_id":       res.Role.ID,
		"role_name":     res.Role.Name,
		"profile_image": res.ProfileImage,
		"token":         res.Token,
	}

	accessToken, _ := pkg.GenerateToken(accessTokenData, configs.TokenSecret, configs.TokenExpired)

	authToken := models.TokenAuth{
		AccessToken: accessToken,
		Type:        "Bearer",
		ExpiredAt:   pkg.CalculateExpiration(time.Now().Add(configs.TokenExpired).Unix()),
	}

	if err != nil {
		switch err.Type {
		case "error_01":
			helpers.ApiResponse(ctx, http.StatusNotFound, "error", "Email not found", nil, nil)
		case "error_02":
			helpers.ApiResponse(ctx, http.StatusInternalServerError, "error", "Internal server error", nil, nil)
		case "error_03":
			helpers.ApiResponse(ctx, http.StatusUnauthorized, "error", "Password is incorrect", nil, nil)
		case "error_04":
			helpers.ApiResponse(ctx, http.StatusInternalServerError, "error", "Internal server error", nil, nil)
		default:
			helpers.ApiResponse(ctx, http.StatusInternalServerError, "error", "Unknown error", nil, nil)
		}
		return
	}

	resData := schemas.LoginUserResponse{
		ID:           res.ID,
		Firstname:    res.Biodata.Firstname,
		Lastname:     res.Biodata.Lastname,
		Email:        res.Biodata.Email,
		RoleId:       res.RoleId,
		RoleName:     res.Role.Name,
		ProfileImage: res.ProfileImage,
		CreatedAt:    res.CreatedAt,
		UpdatedAt:    res.UpdatedAt,
	}
	helpers.ApiResponse(ctx, http.StatusOK, "success", "Login successfully", resData, authToken)
}
