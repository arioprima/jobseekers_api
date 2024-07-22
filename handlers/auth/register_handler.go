package handlers

import (
	"net/http"

	"github.com/arioprima/jobseekers_api/helpers"
	"github.com/arioprima/jobseekers_api/pkg"
	"github.com/arioprima/jobseekers_api/schemas"
	services "github.com/arioprima/jobseekers_api/services/auth"
	"github.com/gin-gonic/gin"
)

type HandlerRegister struct {
	Service services.ServiceRegister
}

func NewHandlerRegister(service services.ServiceRegister) *HandlerRegister {
	return &HandlerRegister{Service: service}
}

func (h *HandlerRegister) RegisterHandler(ctx *gin.Context) {
	var registerRequest schemas.SchemaDataUser
	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err.Error())
		return
	}

	metaConfigs := []schemas.ErrorMetaConfig{
		{
			Tag:     "required",
			Field:   "Firstname",
			Message: "Firstname is required",
		},
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

	errResponse, errCount := pkg.ValidatorRegister(&registerRequest, metaConfigs)
	if errCount > 0 {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, "error", errResponse)
		return
	}

	res, err := h.Service.RegisterService(ctx, nil, &registerRequest)
	if err != nil {
		switch err.Type {
		case "error_01":
			helpers.ValidatorErrorResponse(ctx, http.StatusConflict, "error", "Email already exists")
		case "error_02":
			helpers.ValidatorErrorResponse(ctx, http.StatusInternalServerError, "error", "Internal server error")
		default:
			helpers.ValidatorErrorResponse(ctx, http.StatusInternalServerError, "error", "Internal server error")
		}
		return
	}

	helpers.ApiResponse(ctx, http.StatusCreated, "success", "Register successfully", res, nil)
}
