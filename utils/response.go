package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"website-gin/dto"
	"website-gin/utils/constants"
)

// 通用错误响应
func ResultError(ctx *gin.Context, exception constants.BizExceptionEnum) {
	ctx.JSON(http.StatusOK, dto.Result{
		Success: false,
		Code:    exception.Code,
		Message: exception.Message,
		Data:    nil,
	})
}

// 通用成功响应
func ResultSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, dto.Result{
		Success: true,
		Code:    200,
		Message: "成功",
		Data:    data,
	})
}
