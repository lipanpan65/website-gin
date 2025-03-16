package common

type PageDTO struct {
	PageSize int `form:"page_size,default=10" json:"page_size"`
	Page     int `form:"page,default=1" json:"page"`
}
