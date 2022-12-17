package routes

import (
	"github.com/Safwanseban/Project-Microservices/broker-service/controllers"
	"github.com/Safwanseban/Project-Microservices/broker-service/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(ctx *gin.Engine) {
	ctx.POST("/", middlewares.CORSMiddleware(), controllers.HomeBroker)
	ctx.POST("/handle", middlewares.CORSMiddleware(), controllers.HandleSubmission)

}
