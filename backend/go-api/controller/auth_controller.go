package controller

import (
	"github.com/arioprima/jobseeker/tree/main/backend/go-api/models"
	"github.com/arioprima/jobseeker/tree/main/backend/go-api/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	loginRequest := models.LoginInput{}
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		ctx.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}

	loginResponse, err := controller.AuthService.Login(ctx, loginRequest)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "Internal Server Error",
		})
		return
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"Status":  http.StatusOK,
			"Message": "Success",
			"Data":    loginResponse,
		})
	}
}

func (controller *AuthController) Register(ctx *gin.Context) {
	registerRequest := models.RegisterInput{}
	err := ctx.ShouldBindJSON(&registerRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	registerResponse, err := controller.AuthService.Register(ctx, registerRequest)

	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "Internal Server Error",
		})
	} else {
		ctx.IndentedJSON(http.StatusCreated, gin.H{
			"Status":  http.StatusCreated,
			"Message": "OK",
			"Data":    registerResponse,
		})
	}
}

func (controller *AuthController) VerifyEmail(ctx *gin.Context) {
	verifyRequest := models.VerifyInput{}
	if err := ctx.ShouldBindJSON(&verifyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	verifyResponse, err := controller.AuthService.VerifyEmail(ctx, verifyRequest)
	if err != nil {
		log.Printf("Error verifying email: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Code OTP is invalid", "status": http.StatusInternalServerError})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "OK",
		"data":    verifyResponse,
	})
}
