package repository

import (
	"website-gin/config"
	"website-gin/internal/models"
)

func CreateSubject(subject *models.Subject) error {
	return config.DB.Create(subject).Error
}
