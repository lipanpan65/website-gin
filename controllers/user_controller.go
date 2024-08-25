package controllers

import (
	"github.com/gin-gonic/gin"
	"website-gin/models"
	"website-gin/services"
	"website-gin/utils"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ResponseError(c, 400, "Invalid input")
		return
	}
	createdUser := services.CreateUser(user)
	utils.ResponseSuccess(c, createdUser)
}

func GetAllUsers(c *gin.Context) {
	users := services.GetAllUsers()
	utils.ResponseSuccess(c, users)
}
