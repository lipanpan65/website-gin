package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	v1 "website-gin/api/v1"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {
	apiV1 := r.Group("/api/v1")
	{ // 注册 v1 版本的用户路由
		v1.RegisterUserRoutes(apiV1)
		v1.RegisterSubjectRoutes(apiV1)
		v1.RegisterTopicRoutes(apiV1, db)
	}
	return r
}
