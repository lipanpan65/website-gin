package services

import (
	"website-gin/internal/models"
	"website-gin/internal/repository"
)

func CreateUser(user models.User) models.User {
	return repository.CreateUser(user)
}

func GetAllUsers() []models.User {
	return repository.GetAllUsers()
}
