package services

import (
	"website-gin/dao"
	"website-gin/models"
)

func CreateUser(user models.User) models.User {
	return dao.CreateUser(user)
}

func GetAllUsers() []models.User {
	return dao.GetAllUsers()
}
