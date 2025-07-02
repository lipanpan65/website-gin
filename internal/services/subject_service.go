package services

import (
	"website-gin/internal/model"
	"website-gin/internal/repository"
)

func CreateSubject(subject *model.Subject) error {
	return repository.CreateSubject(subject)
}
