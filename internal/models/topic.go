package models

import "time"

type Topic struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	TopicName string `json:"topic_name" gorm:"column:topic_name"`
	Remark    string `json:"remark" gorm:"column:remark"`
}

type TopicVo struct {
	Topic      `gorm:"embedded"` // GORM 处理嵌套结构体
	CreateTime time.Time         `json:"create_time" gorm:"column:create_time"`
}

// TableName 定义 GORM 表名为 tb_topics
func (Topic) TableName() string {
	return "tb_topics"
}

type PageVo struct {
	PageSize int `form:"page_size,default=10" json:"page_size"`
	Page     int `form:"page,default=1" json:"page"`
}
