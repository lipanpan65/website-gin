package controllers

import (
	"github.com/gin-gonic/gin"
	"website-gin/internal/models"
	"website-gin/internal/services"
	"website-gin/utils"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ResultError(c, nil)
		return
	}
	createdUser := services.CreateUser(user)
	utils.ResultSuccess(c, createdUser)

}

func GetAllUsers(c *gin.Context) {
	users := services.GetAllUsers()
	utils.ResultSuccess(c, users)
}
