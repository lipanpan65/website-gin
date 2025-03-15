package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"website-gin/dto"
	"website-gin/utils/errors"
)

func HandleError(ctx *gin.Context, err interface{}) {
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
