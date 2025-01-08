package services

import (
	"website-gin/models"
	"website-gin/repository"
)

func CreateSubject(subject *models.Subject) error {
	return repository.CreateSubject(subject)
}
