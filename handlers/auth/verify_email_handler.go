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
	//TODO implement me
	userID := ctx.Query("user_id")
	otp := ctx.Query("otp")

	if userID == "" {
		helpers.ValidatorErrorResponse(ctx, 400, "error", "user_id is required")
		return
	} else if otp == "" {
		helpers.ValidatorErrorResponse(ctx, 400, "error", "otp is required")
		return
	}

	res, err := v.Service.VerifyEmailService(ctx, nil, userID, otp)
	if err != nil {
		switch err.Type {
		case "error_01":
			helpers.ValidatorErrorResponse(ctx, 400, "error", "Invalid OTP")
		}
	}

	helpers.ApiResponse(ctx, http.StatusOK, "success", "Register successfully", res, nil)

}
