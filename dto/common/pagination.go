package common

// PageInfo 分页信息
type PageInfo struct {
	Total    int `json:"total"`
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}

// PagedData 分页数据
type PagedData struct {
	Page PageInfo    `json:"page"`
	Data interface{} `json:"data"`
}
