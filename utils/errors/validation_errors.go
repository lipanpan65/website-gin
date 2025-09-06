package errors

// 参数验证错误 (4000-4099)
var (
	InvalidParameter = NewBusinessError("4000", "参数无效")
	RequiredField    = NewBusinessError("4001", "必填字段不能为空")
	InvalidFormat    = NewBusinessError("4002", "字段格式不正确")
	ValueOutOfRange  = NewBusinessError("4003", "字段值超出允许范围")
)
