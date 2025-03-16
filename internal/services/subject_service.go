package services

import (
	"website-gin/internal/models"
	"website-gin/internal/repository"
)

func CreateSubject(subject *models.Subject) error {
	return repository.CreateSubject(subject)
}
