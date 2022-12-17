package routes

import (
	"github.com/Safwanseban/Project-Microservices/logger-service/models"
	"github.com/gin-gonic/gin"
)

type JSONPayload struct {
	Data string `json:"data"`
	Name string `json:"name"`
}

func LogEntry(c *gin.Context) {
	var payload JSONPayload
	var log models.LogEntry
	
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(404, gin.H{
			"err": err.Error(),
		})
	}
	event := models.LogEntry{
		Data: payload.Data,
		Name: payload.Name,
	}
	
	if err := log.Insert(event); err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"msg": "logged",
	})
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Routes(c *gin.Engine) {
	c.POST("/log", CORSMiddleware(), LogEntry)

}
