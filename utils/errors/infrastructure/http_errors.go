package infrastructure

import "website-gin/utils/errors"

// 这里可以定义一些 HTTP 层的错误，例如请求超时、连接错误等
var (
	HTTPRequestTimeout  = errors.NewTechnicalError(504, "HTTP 请求超时")
	HTTPConnectionError = errors.NewTechnicalError(502, "HTTP 连接错误")
)
