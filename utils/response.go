package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"website-gin/dto/response"
	"website-gin/utils/errors"
)

func ResultSuccess(ctx *gin.Context, data interface{}, opt ...interface{}) {
	code, message := parseSuccessOptions(opt...)

	ctx.JSON(http.StatusOK, response.Result{
		Success: true,
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// parseSuccessOptions 解析 ResultSuccess 函数的可选参数
func parseSuccessOptions(opt ...interface{}) (int, string) {
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

	return code, message
}

//func ResultSuccess(ctx *gin.Context, data interface{}, opt ...interface{}) {
//	code := http.StatusOK
//	message := "成功"
//
//	if len(opt) > 0 {
//		switch v := opt[0].(type) {
//		case *errors.BaseError:
//			code = v.Code
//			message = v.Message
//		case string:
//			message = v
//		case int:
//			code = v
//		}
//		if len(opt) > 1 {
//			if v, ok := opt[1].(string); ok {
//				message = v
//			}
//			if v, ok := opt[1].(int); ok {
//				code = v
//			}
//		}
//	}
//
//	ctx.JSON(http.StatusOK, response.Result{
//		Success: true,
//		Code:    code,
//		Message: message,
//		Data:    data,
//	})
//}

// ResultError 返回错误响应
func ResultError(ctx *gin.Context, err interface{}) {
	code, message := parseError(err)

	ctx.JSON(http.StatusOK, response.Result{
		Success: false,
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// parseError 解析错误信息并返回错误码和错误消息
func parseError(err interface{}) (int, string) {
	var code int
	var message string

	switch e := err.(type) {
	case *errors.BaseError:
		code = e.Code
		message = e.Message
	case error:
		code = http.StatusInternalServerError
		message = e.Error()
	case string:
		code = http.StatusBadRequest
		message = e
	default:
		code = http.StatusInternalServerError
		message = "未知错误"
	}

	return code, message
}

//func ResultError(ctx *gin.Context, err interface{}) {
//	var code int
//	var message string
//
//	switch e := err.(type) {
//	case *errors.BaseError:
//		code = e.Code
//		message = e.Message
//	case string:
//		code = http.StatusBadRequest
//		message = e
//	default:
//		code = http.StatusInternalServerError
//		message = "未知错误"
//	}
//
//	ctx.JSON(http.StatusOK, response.Result{
//		Success: false,
//		Code:    code,
//		Message: message,
//		Data:    nil,
//	})
//}
