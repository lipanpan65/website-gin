package errors

// ErrorType 错误类型
type ErrorType string

const (
	TypeBusiness  ErrorType = "business"  // 业务逻辑错误
	TypeDomain    ErrorType = "domain"    // 领域规则错误
	TypeTechnical ErrorType = "technical" // 技术错误（如数据库、网络）
)

// BaseError 基础错误结构体
type BaseError struct {
	Code    int
	Message string
	Type    ErrorType
}

//func (b BaseError) Error() string {
//	//TODO implement me
//	//panic("implement me")
//	return fmt.Sprintf("error type: %s, code: %d, message: %s", b.Type, b.Code, b.Message)
//}

// Error 实现 error 接口
func (be *BaseError) Error() string {
	return be.Message
}

// NewBusinessError 创建业务错误
func NewBusinessError(code int, message string) *BaseError {
	return &BaseError{
		Code:    code,
		Message: message,
		Type:    TypeBusiness,
	}
}

// NewDomainError 创建领域错误
func NewDomainError(code int, message string) *BaseError {
	return &BaseError{
		Code:    code,
		Message: message,
		Type:    TypeDomain,
	}
}

// NewTechnicalError 创建技术错误
func NewTechnicalError(code int, message string) *BaseError {
	return &BaseError{
		Code:    code,
		Message: message,
		Type:    TypeTechnical,
	}
}
