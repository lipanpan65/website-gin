package v1

import (
	"github.com/gin-gonic/gin"
	"website-gin/controllers"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", controllers.CreateUser) // 创建用户
		userGroup.GET("/", controllers.GetAllUsers) // 获取所有用户
	}
}
