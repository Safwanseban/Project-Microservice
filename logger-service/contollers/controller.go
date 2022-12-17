package controllers

// import (
// 	"github.com/Safwanseban/Project-Microservices/logger-service/models"
// 	"github.com/gin-gonic/gin"
// )
// type JSONPayload struct{
// Data string `json:"data"`
// Name string `json:"name"`


// }

// func LogEntry(c *gin.Context) {
// var payload JSONPayload
// // var modelss models.Models
// if err:=c.ShouldBindJSON(&payload);err!=nil{
// c.JSON(404,gin.H{
// 	"err":err.Error(),
// })
// }
// event:=models.LogEntry{
// 	Data: payload.Data,
// 	Name: payload.Name,
// }
// if err:=models.LogEntry.Insert(event);err!=nil{
// 	c.JSON(500,gin.H{
// 		"err":err.Error(),
// 	})
// 	return
// }
// c.JSON(200,gin.H{
// 	"msg":"logged",
// })
// }
