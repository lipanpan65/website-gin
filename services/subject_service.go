package services

import (
	"website-gin/dao"
	"website-gin/models"
)

func CreateSubject(subject *models.Subject) error {
	return dao.CreateSubject(subject)
}
