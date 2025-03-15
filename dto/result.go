package dto

type Result struct {
	Success bool        `json:"success"` // 请求是否成功
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
