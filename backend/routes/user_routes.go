package routes

import (
	"github.com/arioprima/jobseeker/tree/main/backend/controller"
	"github.com/arioprima/jobseeker/tree/main/backend/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(authController *controller.AuthController, adminController *controller.AdminController) *gin.Engine {
	service := gin.Default()
	middleware.SetupCorsMiddleware(service)

	router := service.Group("/api")

	router.POST("/auth/login", authController.Login)
	router.POST("/auth/register", authController.Register)
	router.POST("/auth/verify-email", authController.VerifyEmail)

	router.POST("/admin/save", adminController.Save)
	router.POST("/admin/update", adminController.Update)
	router.POST("/admin/delete", adminController.Delete)
	router.POST("/admin/get", adminController.FindByID)

	return service
}
