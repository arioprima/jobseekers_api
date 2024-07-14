package routes

import (
	handlers "github.com/arioprima/jobseekers_api/handlers/auth"
	repositories "github.com/arioprima/jobseekers_api/repositories/auth"
	services "github.com/arioprima/jobseekers_api/services/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TestMiddleware struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Name2 string `json:"name2"`
}

func SetupAuthRoutes(route *gin.Engine, db *gorm.DB) {
	// Initialize dependencies
	loginRepository := repositories.NewRepositoryLoginImpl(nil, db)
	loginService := services.NewServiceLoginImpl(loginRepository, nil)
	loginHandler := handlers.NewHandlerLogin(loginService)

	// Setup routes
	groupRoute := route.Group("/jobseeker-api/api")
	groupRoute.POST("/login", loginHandler.LoginHandler)

	user := &TestMiddleware{
		ID:    1,
		Name:  uuid.New().String(),
		Name2: uuid.New().String(),
	}

	groupRoute.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, user)
	})

}
