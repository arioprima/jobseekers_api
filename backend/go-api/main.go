package main

import (
	"github.com/arioprima/jobseeker/tree/main/backend/go-api/controller"
	"github.com/arioprima/jobseeker/tree/main/backend/go-api/initializers"
	"github.com/arioprima/jobseeker/tree/main/backend/go-api/repository"
	"github.com/arioprima/jobseeker/tree/main/backend/go-api/routes"
	"github.com/arioprima/jobseeker/tree/main/backend/go-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//checkConeection to database success or not
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Println("Load config error", err)
	}
	db, err := initializers.ConnectDB(&config)
	if err != nil {
		log.Println("Connect to database error", err)
	} else {
		log.Println("Connect to database successfully")
	}
	validate := validator.New()
	authRepository := repository.NewAuthRepositoryImpl(db)
	authService := service.NewAuthServiceImpl(authRepository, db, validate)
	authController := controller.NewAuthController(authService)

	router := routes.UserRouter(authController)
	err = router.Run(":8080")
	if err != nil {
		log.Println("Run router error", err)
	} else {
		log.Println("Run router successfully")
	}
}
