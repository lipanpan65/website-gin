package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"website-gin/dto/response"
	"website-gin/utils/errors"
)

func HandleError(ctx *gin.Context, err interface{}) {
	var code string
	var message string

	switch e := err.(type) {
	case *errors.BaseError:
		code = e.Code
		message = e.Message
	case string:
		code = strconv.Itoa(http.StatusBadRequest)
		message = e
	default:
		code = strconv.Itoa(http.StatusInternalServerError)
		message = "未知错误"
	}

	ctx.JSON(http.StatusOK, response.Result{
		Success: false,
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
