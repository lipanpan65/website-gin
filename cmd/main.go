package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"website-gin/config"
	"website-gin/middleware"
	"website-gin/routes"
)

/**
export GO_ENV=dev
go run cmd/main.go

export GO_ENV=prod
go run cmd/main.go

*/

func main() {
	// 初始化配置
	config.InitConfig()
	log.Printf("Configuration initialized: %+v", config.Conf)
	// 初始化数据库
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	log.Println("Database initialized")
	// 自动迁移数据库
	//if err := db.AutoMigrate(&models.User{}); err != nil {
	//	log.Fatalf("Failed to migrate database: %v", err)
	//}
	//log.Println("Database migrated successfully")

	// 初始化路由
	r := gin.Default()
	// 使用全局异常处理中间件
	r.Use(middleware.GlobalErrorHandler())
	routes.SetupRouter(r, db)
	log.Println("Routes initialized")

	// 使用配置中的 Port 启动服务器
	log.Printf("Starting server on port %s", config.Conf.Port)
	err = r.Run(config.Conf.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
