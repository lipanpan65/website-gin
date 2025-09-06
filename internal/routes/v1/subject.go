package v1

import (
	"github.com/gin-gonic/gin"
	"website-gin/internal/handlers"
)

func RegisterSubjectRoutes(router *gin.RouterGroup) {
	subjectGroup := router.Group("/subject")
	{
		subjectGroup.POST("/", handlers.CreateSubject) // 创建用户
		//userGroup.GET("/", handlers.GetAllUsers) // 获取所有用户
	}
}
