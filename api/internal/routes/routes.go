package router

import (
	"github.com/gin-gonic/gin"
	"github.com/scallyt/neko/internal/handlers"
	"github.com/scallyt/neko/internal/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.GET("/jwt", middlewares.AuthMiddleware, handlers.GetProjectByUserId)

	api.POST("/clerk-webhook", handlers.ClerkWebhookHandler)

	api.POST("/new-project", handlers.CreateProject)

	api.GET("/ping", middlewares.AuthMiddleware, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
