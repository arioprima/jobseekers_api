package handlers

import (
	"github.com/arioprima/jobseekers_api/helpers"
	services "github.com/arioprima/jobseekers_api/services/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VerifyEmailHandler struct {
	Service services.ServiceVerifyEmail
}

func NewVerifyEmailHandler(service services.ServiceVerifyEmail) *VerifyEmailHandler {
	return &VerifyEmailHandler{Service: service}
}

func (v *VerifyEmailHandler) VerifyEmailHandler(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	otp := ctx.Query("otp")

	if userID == "" {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, "error", "user_id is required")
		return
	}
	if otp == "" {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, "error", "otp is required")
		return
	}

	_, err := v.Service.VerifyEmailService(ctx, nil, userID, otp)

	if err != nil {
		switch err.Type {
		case "error_01":
			helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, "error", "Invalid OTP")
		default:
			helpers.ValidatorErrorResponse(ctx, http.StatusInternalServerError, "error", "Internal Server Error")
		}
		return
	}

	helpers.ApiResponse(ctx, http.StatusOK, "success", "Verification successful", nil, nil)
}
