package main

import (
	"context"
	"log"
	"time"

	"github.com/Safwanseban/Project-Microservices/logger-service/configs"
	"github.com/Safwanseban/Project-Microservices/logger-service/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

const mongoURL = "mongodb://localhost:27017"

var R = gin.Default()
var Client *mongo.Client

func main() {

	mongoClient, err := configs.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}
	Client = mongoClient
	// create a ctx inordeer to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	defer func() {
		if err = Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	routes.Routes(R)
	R.Run(":8082")
}
