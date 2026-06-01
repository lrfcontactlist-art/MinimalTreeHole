package main

import (
	"log"

	"github.com/lrfcontactlist-art/MinimalTreeHole/internal/config"
	"github.com/lrfcontactlist-art/MinimalTreeHole/internal/database"
	"github.com/lrfcontactlist-art/MinimalTreeHole/internal/router"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 连接数据库
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Database connected successfully")

	// 设置路由
	r := router.SetupRouter(db)

	// 启动服务器
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
