package v1

import (
	"github.com/gin-gonic/gin"
	"website-gin/controllers"
)

func RegisterTopicRoutes(router *gin.RouterGroup) {
	topicsGroup := router.Group("/topics")
	{
		topicsGroup.GET("/", controllers.QueryTopic)
		topicsGroup.POST("/", controllers.CreateTopic) // 创建用户
	}
}
