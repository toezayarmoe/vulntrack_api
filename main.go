package main

import (
	"github.com/gin-gonic/gin"
	"github.com/toezayarmoe/vulntrack_api/config"
	"github.com/toezayarmoe/vulntrack_api/handler"
	"github.com/toezayarmoe/vulntrack_api/middleware"
	"github.com/toezayarmoe/vulntrack_api/models"
)

func main() {
	config.ConnectDB()
	models.Migrate(config.DB)

	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.POST("/login", handler.Login)
	}

	protected := r.Group("/api")

	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/globalsummary", handler.GlobalSummary)
		protected.GET("/fetchenv", handler.FetchEnvironment)
		protected.GET("/reports", handler.GetReports)
	}
	r.Run(":8080")
}
