package repository

import "website-gin/models"

var users []models.User

func CreateUser(user models.User) models.User {
	user.ID = len(users) + 1
	users = append(users, user)
	return user
}

func GetAllUsers() []models.User {
	return users
}
