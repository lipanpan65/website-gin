package v1

import (
	"github.com/gin-gonic/gin"
	"website-gin/internal/handlers"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", handlers.CreateUser) // 创建用户
		userGroup.GET("/", handlers.GetAllUsers) // 获取所有用户
	}
}
