package routes

import (
	"github.com/arioprima/jobseekers_api/config"
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

func SetupAuthRoutes(route *gin.RouterGroup, db *gorm.DB, cfg config.Config) {
	// Initialize dependencies
	loginRepository := repositories.NewRepositoryLoginImpl(db)
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

	//resend otp
	resendOtpRepository := repositories.NewResendOtpRepositoryImpl(db)
	resendOtpService := services.NewServiceResendOtpImpl(resendOtpRepository)
	resendOtpHandler := handlers.NewResendOtpHandler(resendOtpService)

	// Initialize OAuth
	handlers.InitializeOAuthConfig(cfg)
	// Setup routes
	groupRoute := route.Group("/auth")
	groupRoute.POST("/login", loginHandler.LoginHandler)
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.GET("/verify-email", verifyEmailHandler.VerifyEmailHandler)
	groupRoute.PUT("/resend-otp/:user_id", resendOtpHandler.ResendOtpHandler)
	groupRoute.GET("/login/google", handlers.GoogleLogin)
	groupRoute.GET("/google/callback", handlers.GoogleCallback)

	user := &TestMiddleware{
		ID:    1,
		Name:  uuid.New().String(),
		Name2: uuid.New().String(),
	}

	groupRoute.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, user)
	})

}
