package models

type Topic struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	TopicName string `json:"topic_name" gorm:"column:topic_name"`
	Remark    string `json:"remark" gorm:"column:remark"`
}

// TableName 定义 GORM 表名为 tb_topics
func (Topic) TableName() string {
	return "tb_topics"
}
