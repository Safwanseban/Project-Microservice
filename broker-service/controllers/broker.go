package controllers

import (
	"net/http"

	"github.com/Safwanseban/Project-Microservices/broker-service/helpers"
	"github.com/gin-gonic/gin"
)

func HomeBroker(c *gin.Context) {
	data := helpers.JsonREsponse{
		Error:   false,
		Messege: "hit the broker",
	}
	// if err := c.ShouldBindJSON(&data); err != nil {
	// 	c.JSON(400, gin.H{
	// 		"msg": "error",
	// 		"err": err.Error(),
	// 	})
	// 	c.Abort()
	// 	return
	// }
	// c.IndentedJSON()
	
	c.JSON(http.StatusAccepted, gin.H{
		"msg": data,
	})

}
