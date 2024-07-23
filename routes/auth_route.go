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

func SetupAuthRoutes(route *gin.RouterGroup, db *gorm.DB) {
	// Initialize dependencies
	loginRepository := repositories.NewRepositoryLoginImpl(nil, db)
	loginService := services.NewServiceLoginImpl(loginRepository, nil)
	loginHandler := handlers.NewHandlerLogin(loginService)

	//register
	registerRepository := repositories.NewRegisterRepositoryImpl(db)
	registerService := services.NewServiceRegisterImpl(registerRepository)
	registerHandler := handlers.NewHandlerRegister(registerService)

	//verify email
	verifyEmailRepository := repositories.NewVerifyEmailRepositoryImpl(db)
	verifyEmailService := services.NewServiceVerifyEmailImpl(verifyEmailRepository)
	verifyEmailHandler := handlers.NewVerifyEmailHandler(verifyEmailService)

	// Setup routes
	groupRoute := route.Group("/api")
	groupRoute.POST("/login", loginHandler.LoginHandler)
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.GET("/verify_email", verifyEmailHandler.VerifyEmailHandler)

	user := &TestMiddleware{
		ID:    1,
		Name:  uuid.New().String(),
		Name2: uuid.New().String(),
	}

	groupRoute.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, user)
	})

}
