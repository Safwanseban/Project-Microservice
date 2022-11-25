package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonREsponse struct {
	Error   bool   `json:"error"`
	Messege string `json:"messege"`
	Data    any    `json:"data,omitempty"`
}

func HomeBroker(c *gin.Context) {
	data := JsonREsponse{
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
