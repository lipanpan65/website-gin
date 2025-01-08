package v1

import (
	"github.com/gin-gonic/gin"
	"website-gin/controllers"
)

func RegisterSubjectRoutes(router *gin.RouterGroup) {
	subjectGroup := router.Group("/subject")
	{
		subjectGroup.POST("/", controllers.CreateSubject) // 创建用户
		//userGroup.GET("/", controllers.GetAllUsers) // 获取所有用户
	}
}
