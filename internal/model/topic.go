package model

import "time"

type Topic struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	TopicName string `json:"topic_name" gorm:"column:topic_name"`
	Enable    int    `json:"enable" gorm:"column:enable"`
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
