package main

import (
	"fmt"
	"os"

	"github.com/Safwanseban/Project-Microservices/broker-service/initializers"
	"github.com/Safwanseban/Project-Microservices/broker-service/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.GetEnv()
}

var R = gin.Default()

func main() {
	routes.Routes(R)
	fmt.Printf("PORT starts on %s", os.Getenv("PORT"))
	R.Run()
}
