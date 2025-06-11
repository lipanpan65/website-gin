package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"website-gin/dto/request"
	"website-gin/internal/services"
	"website-gin/utils"
	"website-gin/utils/errors/application"
)

type TopicController struct {
	//topicService *services.TopicService
	topicService services.TopicServiceInterface
}

func NewTopicController(topicService services.TopicServiceInterface) *TopicController {
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
		return
	}
	utils.ResultSuccess(ctx, topicVo)
}

func (c *TopicController) QueryTopicByID(ctx *gin.Context) {
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

func (c *TopicController) QueryTopics(ctx *gin.Context) {
	// 获取分页参数
	pageStr := ctx.DefaultQuery("page", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")

	// 解析分页参数
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		utils.ResultError(ctx, application.ParameterError)
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		utils.ResultError(ctx, application.ParameterError)
		return
	}

	// 获取查询条件
	var conditions = make(map[string]interface{})
	// 这里可以根据实际需求从请求中获取更多的查询条件
	// 例如：category := ctx.Query("category")
	// if category != "" {
	//    conditions["category"] = category
	// }

	// 调用服务层方法进行分页查询
	topics, total, err := c.topicService.QueryTopics(conditions, page, pageSize)
	if err != nil {
		utils.ResultError(ctx, err)
		return
	}

	// 返回带分页信息的成功响应
	utils.ResultSuccessWithPagination(ctx, int(total), page, pageSize, topics)
}
