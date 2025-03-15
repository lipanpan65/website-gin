package constants

// 工单相关的错误
var (
	WorkOrderNotFound     = BizExceptionEnum{400, "工单未找到"}
	WorkOrderAlreadyExist = BizExceptionEnum{400, "工单已存在"}
)
