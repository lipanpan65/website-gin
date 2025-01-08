package routes

import (
	"github.com/gin-gonic/gin"
	v1 "website-gin/api/v1"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//v1 := router.Group("/api/v1")
	//{
	//	v1.POST("/users", controllers.CreateUser)
	//	v1.GET("/users", controllers.GetAllUsers)
	//}
	
	apiV1 := router.Group("/api/v1")
	{
		v1.RegisterUserRoutes(apiV1) // 注册 v1 版本的用户路由
	}
	return router
}
