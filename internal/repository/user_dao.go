package repository

import (
	"website-gin/internal/model"
)

var users []model.User

func CreateUser(user model.User) model.User {
	user.ID = len(users) + 1
	users = append(users, user)
	return user
}

func GetAllUsers() []model.User {
	return users
}
