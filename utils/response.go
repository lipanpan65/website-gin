package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"website-gin/dto"
	"website-gin/utils/errors"
)

func ResultSuccess(ctx *gin.Context, data interface{}, opt ...interface{}) {
	code := http.StatusOK
	message := "成功"

	if len(opt) > 0 {
		switch v := opt[0].(type) {
		case *errors.BaseError:
			code = v.Code
			message = v.Message
		case string:
			message = v
		case int:
			code = v
		}
		if len(opt) > 1 {
			if v, ok := opt[1].(string); ok {
				message = v
			}
			if v, ok := opt[1].(int); ok {
				code = v
			}
		}
	}

	ctx.JSON(http.StatusOK, dto.Result{
		Success: true,
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func ResultError(ctx *gin.Context, err interface{}) {
	var code int
	var message string

	switch e := err.(type) {
	case *errors.BaseError:
		code = e.Code
		message = e.Message
	case string:
		code = http.StatusBadRequest
		message = e
	default:
		code = http.StatusInternalServerError
		message = "未知错误"
	}

	ctx.JSON(http.StatusOK, dto.Result{
		Success: false,
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
