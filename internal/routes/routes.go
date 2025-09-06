package routes

import (
	"github.com/gin-gonic/gin"
	"website-gin/internal/handlers"
	v1 "website-gin/internal/routes/v1"
)

func SetupRouter(r *gin.Engine, topicHandler *handlers.TopicHandler) *gin.Engine {
	apiV1 := r.Group("/api/v1")
	{ // 注册 v1 版本的用户路由
		v1.RegisterUserRoutes(apiV1)
		v1.RegisterSubjectRoutes(apiV1)
		v1.RegisterTopicRoutes(apiV1, topicHandler)
	}
	return r
}
