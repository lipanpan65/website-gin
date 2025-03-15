package constants

// BizExceptionEnum 业务异常枚举
type BizExceptionEnum struct {
	Code    int
	Message string
}

// 定义常见的错误枚举
var (
	DictExisted         = BizExceptionEnum{400, "字典已经存在"}
	ErrorCreateDict     = BizExceptionEnum{500, "创建字典失败"}
	ErrorWrapperField   = BizExceptionEnum{500, "包装字典属性失败"}
	FileReadingError    = BizExceptionEnum{400, "FILE_READING_ERROR!"}
	FileNotFound        = BizExceptionEnum{400, "FILE_NOT_FOUND!"}
	UploadError         = BizExceptionEnum{500, "上传图片出错"}
	DbResourceNull      = BizExceptionEnum{400, "数据库中没有该资源"}
	NoPermission        = BizExceptionEnum{405, "权限异常"}
	NoAccess            = BizExceptionEnum{403, "没有权限"}
	RequestInvalidate   = BizExceptionEnum{400, "请求数据格式不正确"}
	InvalidCaptcha      = BizExceptionEnum{400, "验证码不正确"}
	CantDeleteAdmin     = BizExceptionEnum{600, "不能删除超级管理员"}
	CantFreezeAdmin     = BizExceptionEnum{600, "不能冻结超级管理员"}
	CantChangeAdmin     = BizExceptionEnum{600, "不能修改超级管理员角色"}
	UserAlreadyReg      = BizExceptionEnum{401, "该用户已经注册"}
	NoThisUser          = BizExceptionEnum{400, "没有此用户"}
	UserNotExisted      = BizExceptionEnum{400, "没有此用户"}
	AccountFrozen       = BizExceptionEnum{401, "账号被冻结"}
	OldPwdNotRight      = BizExceptionEnum{402, "原密码不正确"}
	TwoPwdNotMatch      = BizExceptionEnum{405, "两次输入密码不一致"}
	RoomNotExisted      = BizExceptionEnum{400, "Room don't existed."}
	RuleNotExisted      = BizExceptionEnum{400, "Rule config don't existed."}
	OrderCanNotConfirm  = BizExceptionEnum{400, "Current order cannot be confirmed."}
	MenuCodeCoincidence = BizExceptionEnum{400, "菜单编号和副编号不能一致"}
	ExistedTheMenu      = BizExceptionEnum{400, "菜单编号重复，不能添加"}
	DictMustBeNumber    = BizExceptionEnum{400, "字典的值必须为数字"}
	DeptDuplicateName   = BizExceptionEnum{400, "该职场已存在"}
	RequestNull         = BizExceptionEnum{400, "请求有错误"}
	SessionTimeout      = BizExceptionEnum{400, "会话超时"}
	DatabaseError       = BizExceptionEnum{500, "数据库异常"}
	ParameterError      = BizExceptionEnum{400, "Parameter error."}
	ServerError         = BizExceptionEnum{500, "服务器异常"}
)
