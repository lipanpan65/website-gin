package dao

import (
	"website-gin/config"
	"website-gin/models"
)

func CreateSubject(subject *models.Subject) error {
	return config.DB.Create(subject).Error
}
