package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var collection *mongo.Collection

func ConnectToMongo() (*mongo.Client, error) {

	clientOptions := option.Client().ApplyURI("mongodb://localhost:27017")

	username := os.Getenv("MONGO_DB_USERNAME")
	username := os.Getenv("MONGO_DB_PASSWORD")

	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to mongo")

	return client, nil
}
