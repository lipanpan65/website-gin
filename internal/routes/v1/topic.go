package v1

import (
	"github.com/gin-gonic/gin"
	"website-gin/internal/handlers"
)

func RegisterTopicRoutes(group *gin.RouterGroup, topicHandler *handlers.TopicHandler) {
	topicsGroup := group.Group("/topics")
	{
		topicsGroup.GET("/", topicHandler.QueryTopics)
		topicsGroup.POST("/", topicHandler.CreateTopic)
	}
}
