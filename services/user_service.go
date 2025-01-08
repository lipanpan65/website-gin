package services

import (
	"website-gin/models"
	"website-gin/repository"
)

func CreateUser(user models.User) models.User {
	return repository.CreateUser(user)
}

func GetAllUsers() []models.User {
	return repository.GetAllUsers()
}
