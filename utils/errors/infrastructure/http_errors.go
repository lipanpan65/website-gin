package infrastructure

import "website-gin/utils/errors"

var (
	HTTPRequestTimeout  = errors.NewTechnicalError("5040", "HTTP 请求超时")
	HTTPConnectionError = errors.NewTechnicalError("5020", "HTTP 连接错误")
)
