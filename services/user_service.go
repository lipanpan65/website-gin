package services

import (
	"website-gin/models"
	"website-gin/repositories"
)

func CreateUser(user models.User) models.User {
	return repositories.CreateUser(user)
}

func GetAllUsers() []models.User {
	return repositories.GetAllUsers()
}
