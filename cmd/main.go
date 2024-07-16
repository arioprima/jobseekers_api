package main

import (
	"github.com/arioprima/jobseekers_api/config"
	"github.com/arioprima/jobseekers_api/middlewares"
	"github.com/arioprima/jobseekers_api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	log := config.NewLogger()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	db, err := config.OpenConnection(&loadConfig)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	router := gin.Default()

	middlewares.SetupCorsMiddleware(router)

	routeGroup := router.Group("/job-vacancies-api")

	routes.SetupAuthRoutes(routeGroup, db)

	port := loadConfig.PORT
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
