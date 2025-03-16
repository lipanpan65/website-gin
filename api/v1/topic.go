package v1

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"website-gin/controllers"
	"website-gin/internal/repository"
	"website-gin/internal/services"
)

func RegisterTopicRoutes(group *gin.RouterGroup, db *gorm.DB) {
	topicRepo := repository.NewTopicRepository(db)
	topicService := services.NewTopicService(topicRepo)
	topicController := controllers.NewTopicController(topicService)
	topicsGroup := group.Group("/topics")
	{
		topicsGroup.GET("/", topicController.QueryTopics)
		topicsGroup.POST("/", topicController.CreateTopic) // 创建用户
	}
}
