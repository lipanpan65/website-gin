package errors

import "fmt"

// ErrorType 定义错误类型枚举
type ErrorType string

const (
	TypeBusiness  ErrorType = "business"  // 业务逻辑错误
	TypeTechnical ErrorType = "technical" // 技术错误（基础设施、外部服务等）
)

// BaseError 基础错误结构体，实现了 error 接口
// 提供统一的错误处理机制，包含错误码、消息和类型信息
type BaseError struct {
	Code    string    // 错误码，用于程序处理和错误识别
	Message string    // 错误消息，用于用户展示
	Type    ErrorType // 错误类型，用于分类处理
}

// Error 实现 error 接口，返回错误消息
func (be *BaseError) Error() string {
	return be.Message
}

// GetCode 获取错误码
func (be *BaseError) GetCode() string {
	return be.Code
}

// GetType 获取错误类型
func (be *BaseError) GetType() ErrorType {
	return be.Type
}

// String 返回格式化的错误信息，包含完整的错误详情
// 格式: [错误类型] 错误码: 错误消息
func (be *BaseError) String() string {
	return fmt.Sprintf("[%s] %s: %s", be.Type, be.Code, be.Message)
}

// IsBusinessError 判断是否为业务错误
func (be *BaseError) IsBusinessError() bool {
	return be.Type == TypeBusiness
}

// IsTechnicalError 判断是否为技术错误
func (be *BaseError) IsTechnicalError() bool {
	return be.Type == TypeTechnical
}

// NewBusinessError 创建业务逻辑错误
// 业务错误通常由用户操作、业务规则违反等引起
//
// 参数:
//   - code: 业务错误码，建议使用 1000-4999 区间
//   - message: 用户可读的错误消息
//
// 返回值:
//   - *BaseError: 业务错误实例
//
// 示例:
//   err := NewBusinessError("1001", "用户不存在")
func NewBusinessError(code string, message string) *BaseError {
	return &BaseError{
		Code:    code,
		Message: message,
		Type:    TypeBusiness,
	}
}

// NewTechnicalError 创建技术错误
// 技术错误通常由基础设施故障、外部服务异常等引起
//
// 参数:
//   - code: 技术错误码，建议使用 5000-5999 区间
//   - message: 技术相关的错误消息
//
// 返回值:
//   - *BaseError: 技术错误实例
//
// 示例:
//   err := NewTechnicalError("5000", "数据库连接失败")
func NewTechnicalError(code string, message string) *BaseError {
	return &BaseError{
		Code:    code,
		Message: message,
		Type:    TypeTechnical,
	}
}
