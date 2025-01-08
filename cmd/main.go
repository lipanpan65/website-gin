package main

import (
	"website-gin/config"
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
	// 初始化数据库
	config.InitDB()
	// 自动迁移数据库
	// config.DB.AutoMigrate(&models.User{})
	// 初始化路由
	router := routes.SetupRouter()
	// 使用配置中的 Port 启动服务器
	err := router.Run(config.Conf.Port)
	if err != nil {
		return
	}

}
