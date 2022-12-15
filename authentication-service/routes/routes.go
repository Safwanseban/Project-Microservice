package routes

import (
	"github.com/Safwanseban/Project-Microservices/authentication-service/controllers"
	"github.com/Safwanseban/Project-Microservices/authentication-service/middlewares"
	"github.com/gin-gonic/gin"
)

func Routes(ctx *gin.Engine) {
	ctx.POST("/", middlewares.CORSMiddleware(), controllers.HomeBroker)
}
