package request

import "github.com/go-playground/validator/v10"

type TopicDTO struct {
	TopicName string `json:"topic_name" gorm:"column:topic_name"`
	Enable    int    `json:"enable" gorm:"column:enable"`
	Remark    string `json:"remark" gorm:"column:remark"`
}

// Validate 验证 TopicDTO 数据
func (u *TopicDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
