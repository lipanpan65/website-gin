package controllers

import (
	"github.com/gin-gonic/gin"
	"website-gin/dto/request"
	"website-gin/internal/services"
	"website-gin/utils"
	"website-gin/utils/errors/application"
)

type TopicController struct {
	topicService *services.TopicService
}

func NewTopicController(topicService *services.TopicService) *TopicController {
	return &TopicController{
		topicService: topicService,
	}
}

func (c *TopicController) CreateTopic(ctx *gin.Context) {
	// 解析客户端请求（URL 参数、Query 参数、Body 数据等）。
	var topicDTO request.TopicDTO
	if err := ctx.ShouldBind(&topicDTO); err != nil {
		utils.ResultError(ctx, application.CantChangeAdmin)
		return
	}
	topicVo, err := c.topicService.CreateTopic(&topicDTO)
	if err != nil {
		utils.HandleError(ctx, err)
	}
	utils.ResultSuccess(ctx, topicVo)
}

func (c *TopicController) QueryTopic(ctx *gin.Context) {
	var id uint
	if err := ctx.ShouldBindUri(&struct {
		ID uint `uri:"id" binding:"required"`
	}{id}); err != nil {
		utils.ResultError(ctx, err.Error())
		return
	}
	topicVo, err := c.topicService.QueryTopicByID(id)
	if err != nil {
		utils.ResultError(ctx, err.Error())
		return
	}
	utils.ResultSuccess(ctx, topicVo)
}
