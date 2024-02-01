package controllers

import (
	"kitten/models"
	"kitten/utils"

	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.JSON(400, gin.H{"error": "User Already Exists"})
		return
	}

	var errHash error

	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "Could not generate password hash"})
		return
	}
	models.DB.Create(&user)

	c.JSON(200, gin.H{"success": "user created"})
}
