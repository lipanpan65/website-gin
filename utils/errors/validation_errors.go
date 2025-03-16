package errors

var (
	InvalidParameter = NewBusinessError(400, "参数无效")
)
