package request

import "github.com/go-playground/validator/v10"

type TopicDTO struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

// Validate 验证 TopicDTO 数据
func (u *TopicDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
