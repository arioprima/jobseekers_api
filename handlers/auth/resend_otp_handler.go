package handlers

import (
	"github.com/arioprima/jobseekers_api/helpers"
	"github.com/arioprima/jobseekers_api/schemas"
	services "github.com/arioprima/jobseekers_api/services/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResendOtpHandler struct {
	Service services.ServiceResendOtp
}

func NewResendOtpHandler(service services.ServiceResendOtp) *ResendOtpHandler {
	return &ResendOtpHandler{Service: service}
}

func (r *ResendOtpHandler) ResendOtpHandler(ctx *gin.Context) {
	var userId = ctx.Param("user_id")
	var resendOtpRequest schemas.SchemaDataUser
	if err := ctx.ShouldBindJSON(&resendOtpRequest); err != nil {
		helpers.ValidatorErrorResponse(ctx, http.StatusBadRequest, "error", err.Error())
		return
	}

	_, err := r.Service.ResendOtp(ctx, &resendOtpRequest, userId)
	if err != nil {
		switch err.Type {
		case "error_01":
			helpers.ValidatorErrorResponse(ctx, http.StatusNotFound, "error", "User not found")
		default:
			helpers.ValidatorErrorResponse(ctx, http.StatusInternalServerError, "error", "Internal Server Error")
		}
		return
	}

	helpers.ApiResponse(ctx, http.StatusOK, "success", "Resend OTP successfully", nil, nil)
}
