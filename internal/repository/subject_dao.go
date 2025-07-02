package repository

import (
	"website-gin/config"
	"website-gin/internal/model"
)

func CreateSubject(subject *model.Subject) error {
	return config.DB.Create(subject).Error
}
