package errors

// ===========================================
// 业务错误 - 业务逻辑相关错误
// ===========================================

// 通用业务错误 (1000-1099)
var (
	DataExisted       = NewBusinessError("1000", "数据已存在，请勿重复创建")
	DbResourceNull    = NewBusinessError("1001", "数据库中没有该资源")
	RequestInvalidate = NewBusinessError("1002", "请求数据格式不正确")
	RequestNull       = NewBusinessError("1003", "请求有错误")
	ParameterError    = NewBusinessError("1004", "Parameter error.")
)

// 用户相关错误 (1100-1199)
var (
	UserAlreadyReg  = NewBusinessError("1100", "该用户已经注册")
	NoThisUser      = NewBusinessError("1101", "没有此用户")
	UserNotExisted  = NewBusinessError("1102", "没有此用户")
	AccountFrozen   = NewBusinessError("1103", "账号被冻结")
	OldPwdNotRight  = NewBusinessError("1104", "原密码不正确")
	TwoPwdNotMatch  = NewBusinessError("1105", "两次输入密码不一致")
)

// 权限相关错误 (1200-1299)
var (
	NoPermission    = NewBusinessError("1200", "权限异常")
	NoAccess        = NewBusinessError("1201", "没有权限")
	CantDeleteAdmin = NewBusinessError("1202", "不能删除超级管理员")
	CantFreezeAdmin = NewBusinessError("1203", "不能冻结超级管理员")
	CantChangeAdmin = NewBusinessError("1204", "不能修改超级管理员角色")
)

// 字典相关错误 (1300-1399)
var (
	DictExisted        = NewBusinessError("1300", "字典已经存在")
	ErrorCreateDict    = NewBusinessError("1301", "创建字典失败")
	ErrorWrapperField  = NewBusinessError("1302", "包装字典属性失败")
	DictMustBeNumber   = NewBusinessError("1303", "字典的值必须为数字")
)

// 文件相关错误 (1400-1499)
var (
	FileReadingError = NewBusinessError("1400", "FILE_READING_ERROR!")
	FileNotFound     = NewBusinessError("1401", "FILE_NOT_FOUND!")
	UploadError      = NewBusinessError("1402", "上传图片出错")
)

// 菜单相关错误 (1500-1599)
var (
	MenuCodeCoincidence = NewBusinessError("1500", "菜单编号和副编号不能一致")
	ExistedTheMenu      = NewBusinessError("1501", "菜单编号重复，不能添加")
)

// 验证相关错误 (1600-1699)
var (
	InvalidCaptcha = NewBusinessError("1600", "验证码不正确")
)

// 组织相关错误 (1700-1799)
var (
	DeptDuplicateName = NewBusinessError("1700", "该职场已存在")
)

// 会话相关错误 (1800-1899)
var (
	SessionTimeout = NewBusinessError("1800", "会话超时")
)
