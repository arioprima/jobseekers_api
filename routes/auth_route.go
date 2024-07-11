package routes

import (
	handlers "github.com/arioprima/jobseekers_api/handlers/auth"
	repositories "github.com/arioprima/jobseekers_api/repositories/auth"
	services "github.com/arioprima/jobseekers_api/services/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoutes(route *gin.Engine, db *gorm.DB) {
	// Initialize dependencies
	loginRepository := repositories.NewRepositoryLoginImpl(nil, db)
	loginService := services.NewServiceLoginImpl(loginRepository, nil)
	loginHandler := handlers.NewHandlerLogin(loginService)

	// Setup routes
	groupRoute := route.Group("/api")
	groupRoute.POST("/login", loginHandler.LoginHandler)
}
