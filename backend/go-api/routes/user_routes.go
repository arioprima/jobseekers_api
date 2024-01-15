package routes

import (
	"github.com/arioprima/Jobseeker/tree/main/backend/go-api/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(authController *controller.AuthController) *gin.Engine {
	service := gin.Default()

	router := service.Group("/api/auth")

	router.POST("/login", authController.Login)
	router.POST("/register", authController.Register)
	router.POST("/verify-email", authController.VerifyEmail)

	return service
}
