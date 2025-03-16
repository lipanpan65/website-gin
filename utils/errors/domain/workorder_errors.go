package domain

import "website-gin/utils/errors"

// 工单领域相关错误枚举
var (
	RoomNotExisted         = errors.NewDomainError("400", "Room don't existed.")
	RuleNotExisted         = errors.NewDomainError("400", "Rule config don't existed.")
	OrderCanNotConfirm     = errors.NewDomainError("400", "Current order cannot be confirmed.")
	WorkOrderNotFound      = errors.NewDomainError("404", "工单未找到")
	WorkOrderInvalidStatus = errors.NewDomainError("400", "工单状态非法")
)
