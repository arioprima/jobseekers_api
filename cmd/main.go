package main

import (
	"github.com/arioprima/jobseekers_api/config"
	_ "github.com/arioprima/jobseekers_api/docs"
	"github.com/arioprima/jobseekers_api/middlewares"
	"github.com/arioprima/jobseekers_api/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	store := cookie.NewStore([]byte(loadConfig.TokenSecret))
	router.Use(sessions.Sessions("jobvacancies_seasion", store))

	middlewares.SetupCorsMiddleware(router)

	router.GET("/swagger.yaml", func(c *gin.Context) {
		c.File("docs/swagger.yaml")
	})

	router.GET("/job-vacancies-api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger.yaml")))

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"code":    "PAGE_NOT_FOUND",
			"message": "Page not found",
		})
	})

	routeGroup := router.Group("/job-vacancies-api")

	routes.SetupAuthRoutes(routeGroup, db, loadConfig)

	port := loadConfig.PORT
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
