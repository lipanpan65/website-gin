package errors

// 工单相关错误枚举 (2000-2099)
var (
	RoomNotExisted         = NewBusinessError("2000", "Room don't existed.")
	RuleNotExisted         = NewBusinessError("2001", "Rule config don't existed.")
	OrderCanNotConfirm     = NewBusinessError("2002", "Current order cannot be confirmed.")
	WorkOrderNotFound      = NewBusinessError("2003", "工单未找到")
	WorkOrderInvalidStatus = NewBusinessError("2004", "工单状态非法")
)
