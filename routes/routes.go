package routes

import (
	"github.com/gin-gonic/gin"
	v1 "website-gin/api/v1"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	{ // 注册 v1 版本的用户路由
		v1.RegisterUserRoutes(apiV1)
		v1.RegisterSubjectRoutes(apiV1)
		v1.RegisterTopicRoutes(apiV1)
	}
	return router
}
