package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"website-gin/dto/common"
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

// ResultSuccessWithPagination 返回带分页信息的成功响应
func ResultSuccessWithPagination(ctx *gin.Context, total int, current int, pageSize int, data interface{}, opt ...interface{}) {
	code, message := parseSuccessOptions(opt...)

	pageInfo := common.PageInfo{
		Total:    total,
		Current:  current,
		PageSize: pageSize,
	}

	pagedData := common.PagedData{
		Page: pageInfo,
		Data: data,
	}

	ctx.JSON(http.StatusOK, response.Result{
		Success: true,
		Code:    code,
		Message: message,
		Data:    pagedData,
	})
}

// parseSuccessOptions 解析 ResultSuccess 函数的可选参数
func parseSuccessOptions(opt ...interface{}) (string, string) {
	// 初始化 code 和 message 的默认值
	code := "0000"
	message := "操作成功"

	if len(opt) > 0 {
		switch v := opt[0].(type) {
		case *errors.BaseError:
			// 从 BaseError 结构体中获取 code 和 message
			code = v.Code
			message = v.Message
		case string:
			// 如果传入的是字符串，将其作为 message
			message = v
		case int:
			// 如果传入的是整数，将其转换为字符串作为 code
			code = fmt.Sprintf("%d", v)
		}

		if len(opt) > 1 {
			if v, ok := opt[1].(string); ok {
				// 如果第二个参数是字符串，更新 message
				message = v
			}
			if v, ok := opt[1].(int); ok {
				// 如果第二个参数是整数，将其转换为字符串更新 code
				code = fmt.Sprintf("%d", v)
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
func parseError(err interface{}) (string, string) {
	var code string = "9999"
	var message string

	switch e := err.(type) {
	case *errors.BaseError:
		code = e.Code
		message = e.Message
	case error:
		code = fmt.Sprintf("%d", http.StatusInternalServerError)
		message = e.Error()
	case string:
		code = fmt.Sprintf("%d", http.StatusBadRequest)
		message = e
	default:
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
