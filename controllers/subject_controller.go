package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"website-gin/models"
	"website-gin/services"
	"website-gin/utils"
)

func CreateSubject(c *gin.Context) {
	var subject models.Subject
	// 绑定 JSON 数据到模型
	if err := c.ShouldBindJSON(&subject); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	// 调用服务层创建用户
	if err := services.CreateSubject(&subject); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 返回成功的响应
	utils.ResponseSuccess(c, subject)
}
