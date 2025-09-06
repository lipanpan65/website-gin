package handlers

import (
	"github.com/gin-gonic/gin"
	"website-gin/internal/models"
	"website-gin/internal/services"
	"website-gin/utils"
	"website-gin/utils/errors"
)

func CreateSubject(c *gin.Context) {
	var subject models.Subject
	// 绑定 JSON 数据到模型
	if err := c.ShouldBindJSON(&subject); err != nil {
		utils.ResultError(c, errors.AccountFrozen)
		return
	}

	// 调用服务层创建用户
	if err := services.CreateSubject(&subject); err != nil {
		utils.ResultError(c, errors.ErrorCreateDict)
		return
	}

	// 返回成功的响应
	utils.ResultSuccess(c, subject)
}
