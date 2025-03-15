package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"website-gin/internal/models"
	"website-gin/internal/repository"
)

func CreateTopic(ctx *gin.Context) {
	// 解析客户端请求（URL 参数、Query 参数、Body 数据等）。
	var topic models.Topic
	if err := ctx.ShouldBind(&topic); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// 创建 topic
	createTopic, err := repository.CreateTopic(topic)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{
		"message": "success",
		"topic":   createTopic,
	})
}

func QueryTopic(ctx *gin.Context) {

}
