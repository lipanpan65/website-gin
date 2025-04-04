package application

import "website-gin/utils/errors"

// 业务逻辑相关错误枚举
var (
	DictExisted         = errors.NewBusinessError("4000", "字典已经存在")
	ErrorCreateDict     = errors.NewBusinessError("500", "创建字典失败")
	ErrorWrapperField   = errors.NewBusinessError("500", "包装字典属性失败")
	FileReadingError    = errors.NewBusinessError("400", "FILE_READING_ERROR!")
	FileNotFound        = errors.NewBusinessError("400", "FILE_NOT_FOUND!")
	UploadError         = errors.NewBusinessError("500", "上传图片出错")
	DbResourceNull      = errors.NewBusinessError("400", "数据库中没有该资源")
	NoPermission        = errors.NewBusinessError("405", "权限异常")
	NoAccess            = errors.NewBusinessError("403", "没有权限")
	RequestInvalidate   = errors.NewBusinessError("400", "请求数据格式不正确")
	InvalidCaptcha      = errors.NewBusinessError("400", "验证码不正确")
	CantDeleteAdmin     = errors.NewBusinessError("600", "不能删除超级管理员")
	CantFreezeAdmin     = errors.NewBusinessError("600", "不能冻结超级管理员")
	CantChangeAdmin     = errors.NewBusinessError("600", "不能修改超级管理员角色")
	UserAlreadyReg      = errors.NewBusinessError("401", "该用户已经注册")
	NoThisUser          = errors.NewBusinessError("400", "没有此用户")
	UserNotExisted      = errors.NewBusinessError("400", "没有此用户")
	AccountFrozen       = errors.NewBusinessError("401", "账号被冻结")
	OldPwdNotRight      = errors.NewBusinessError("402", "原密码不正确")
	TwoPwdNotMatch      = errors.NewBusinessError("405", "两次输入密码不一致")
	MenuCodeCoincidence = errors.NewBusinessError("400", "菜单编号和副编号不能一致")
	ExistedTheMenu      = errors.NewBusinessError("400", "菜单编号重复，不能添加")
	DictMustBeNumber    = errors.NewBusinessError("400", "字典的值必须为数字")
	DeptDuplicateName   = errors.NewBusinessError("400", "该职场已存在")
	RequestNull         = errors.NewBusinessError("400", "请求有错误")
	SessionTimeout      = errors.NewBusinessError("400", "会话超时")
	ParameterError      = errors.NewBusinessError("400", "Parameter error.")
)
