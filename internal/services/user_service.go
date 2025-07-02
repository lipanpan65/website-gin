package services

import (
	"website-gin/internal/model"
	"website-gin/internal/repository"
)

func CreateUser(user model.User) model.User {
	return repository.CreateUser(user)
}

func GetAllUsers() []model.User {
	return repository.GetAllUsers()
}
