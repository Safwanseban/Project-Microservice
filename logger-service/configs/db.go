package configs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoURL = "mongodb://localhost:27017"

var Client *mongo.Client

func ConnectToMongo() (*mongo.Client, error) {
	var err error
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("eror connecting", err)
		return nil, err
	}
	fmt.Println("connected to mongo")
	return Client, nil
}
