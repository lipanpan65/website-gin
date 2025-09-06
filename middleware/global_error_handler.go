package middleware

import (
	stdErrors "errors" // 为标准库的 errors 包设置别名
	"github.com/gin-gonic/gin"
	"website-gin/utils"
	customErrors "website-gin/utils/errors" // 为自定义的 errors 包设置别名
)

// GlobalErrorHandler 全局异常处理中间件
func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var err error
				switch v := r.(type) {
				case error:
					err = v
				case string:
					err = customErrors.NewTechnicalError("500", v)
				default:
					err = customErrors.NewTechnicalError("9999", "未知错误")
				}

				var baseErr *customErrors.BaseError
				// 使用标准库 errors 的 As 方法来处理包装错误
				if stdErrors.As(err, &baseErr) {
					utils.ResultError(c, baseErr)
				} else {
					utils.ResultError(c, customErrors.HTTPRequestTimeout)
				}
				c.Abort()
			}
		}()
		c.Next()
	}
}
