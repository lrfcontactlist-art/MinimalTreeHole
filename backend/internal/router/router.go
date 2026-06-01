package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/lrfcontactlist-art/MinimalTreeHole/internal/handler"
	"github.com/lrfcontactlist-art/MinimalTreeHole/internal/middleware"
	"github.com/lrfcontactlist-art/MinimalTreeHole/internal/repository"
	"github.com/lrfcontactlist-art/MinimalTreeHole/internal/service"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// CORS 中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// 安全中间件
	r.Use(middleware.RateLimitMiddleware())
	r.Use(middleware.SanitizeMiddleware())

	// 初始化依赖
	messageRepo := repository.NewMessageRepository(db)
	messageService := service.NewMessageService(messageRepo)
	messageHandler := handler.NewMessageHandler(messageService)

	// API 路由
	api := r.Group("/api")
	{
		api.GET("/health", messageHandler.HealthCheck)
		api.POST("/messages", messageHandler.CreateMessage)
		api.GET("/messages", messageHandler.GetMessages)
		api.POST("/messages/:id/hug", messageHandler.HugMessage)
	}

	return r
}
