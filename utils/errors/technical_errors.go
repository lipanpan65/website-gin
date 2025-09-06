package errors

// ===========================================
// 技术错误 - 基础设施和外部服务相关错误
// ===========================================

// 数据库相关错误 (5000-5009)
var (
	DatabaseError         = NewTechnicalError("5000", "数据库异常")
	DatabaseTimeout       = NewTechnicalError("5001", "数据库操作超时")
	DatabaseConnectionErr = NewTechnicalError("5002", "数据库连接失败")
	DatabaseQueryError    = NewTechnicalError("5003", "数据库查询错误")
)

// 缓存相关错误 (5010-5019)
var (
	CacheConnectionError = NewTechnicalError("5010", "Redis连接错误")
	CacheTimeout        = NewTechnicalError("5011", "缓存操作超时")
	CacheKeyNotFound    = NewTechnicalError("5012", "缓存键不存在")
	CacheOperationError = NewTechnicalError("5013", "缓存操作失败")
)

// 消息队列相关错误 (5020-5029)
var (
	MQConnectionError = NewTechnicalError("5020", "消息队列连接错误")
	MQPublishFailed  = NewTechnicalError("5021", "消息发布失败")
	MQConsumeFailed  = NewTechnicalError("5022", "消息消费失败")
	MQTimeout        = NewTechnicalError("5023", "消息队列操作超时")
)

// 文件系统相关错误 (5030-5039)
var (
	FileSystemError       = NewTechnicalError("5030", "文件系统错误")
	DiskSpaceInsufficient = NewTechnicalError("5031", "磁盘空间不足")
	FilePermissionDenied  = NewTechnicalError("5032", "文件权限被拒绝")
	FileNotAccessible     = NewTechnicalError("5033", "文件无法访问")
)

// HTTP/网络相关错误 (5040-5049)
var (
	HTTPRequestTimeout  = NewTechnicalError("5040", "HTTP请求超时")
	HTTPConnectionError = NewTechnicalError("5041", "HTTP连接错误")
	NetworkTimeout      = NewTechnicalError("5042", "网络超时")
	NetworkUnavailable  = NewTechnicalError("5043", "网络不可用")
	DNSResolutionError  = NewTechnicalError("5044", "DNS解析失败")
)

// 配置相关错误 (5050-5059)
var (
	ConfigurationError = NewTechnicalError("5050", "配置错误")
	EnvironmentError   = NewTechnicalError("5051", "环境变量错误")
	ConfigNotFound     = NewTechnicalError("5052", "配置文件未找到")
	ConfigParseError   = NewTechnicalError("5053", "配置解析错误")
)

// 第三方服务错误 (5100-5199)
var (
	ThirdPartyAPIError  = NewTechnicalError("5100", "第三方API调用失败")
	PaymentServiceError = NewTechnicalError("5110", "支付服务异常")
	SMSServiceError     = NewTechnicalError("5120", "短信服务异常")
	EmailServiceError   = NewTechnicalError("5130", "邮件服务异常")
	OSSServiceError     = NewTechnicalError("5140", "对象存储服务异常")
	CDNServiceError     = NewTechnicalError("5150", "CDN服务异常")
)

// 系统资源相关错误 (5200-5209)
var (
	SystemResourceError = NewTechnicalError("5200", "系统资源不足")
	MemoryInsufficient = NewTechnicalError("5201", "内存不足")
	CPUOverload        = NewTechnicalError("5202", "CPU负载过高")
	ThreadPoolExhausted = NewTechnicalError("5203", "线程池耗尽")
)

// 安全相关技术错误 (5300-5309)
var (
	EncryptionError    = NewTechnicalError("5300", "加密操作失败")
	DecryptionError    = NewTechnicalError("5301", "解密操作失败")
	CertificateError   = NewTechnicalError("5302", "证书验证失败")
	TokenGenerateError = NewTechnicalError("5303", "Token生成失败")
)