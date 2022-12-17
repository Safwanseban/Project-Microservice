package controllers

import (
	"fmt"

	"github.com/Safwanseban/Project-Microservices/authentication-service/configs"
	"github.com/Safwanseban/Project-Microservices/authentication-service/models"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.GetEnv()
}
func Authenticate(c *gin.Context) {
	var user models.User
	var Autheticate struct {
		Email    string
		Password string
	}
	if err := c.ShouldBindJSON(&Autheticate); err != nil {
		c.JSON(404, gin.H{
			"err": err.Error(),
		})
		return
	}

	users, err := user.GetByEmail(Autheticate.Email)
	if err != nil {
		c.JSON(404, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err := user.HashPassword(Autheticate.Password); err != nil {
		c.JSON(403, gin.H{
			"err": err.Error(),
		})
		return
	}

	fmt.Println(Autheticate.Password)
	// if err := user.CheckPassword(Autheticate.Password); err != nil {
	// 	c.JSON(403, gin.H{
	// 		"err": err.Error(),
	// 	})
	// 	return
	// }

	if err := user.CheckPassword(Autheticate.Password); err != nil {
		c.JSON(403, gin.H{
			"err": err.Error(),
		})
		return
	}

	// valid, err := user.PasswordMatches(Autheticate.Password)
	// if err != nil || !valid {
	// 	c.JSON(403, gin.H{
	// 		"err": err.Error(),
	// 	})
	// 	return
	// }
	c.JSON(200, gin.H{
		"status": true,
		"msg":    users.Email,
	})
}
