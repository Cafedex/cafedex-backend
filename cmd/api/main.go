package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/cafedex-backend/db"
	"github.com/cafedex-backend/handlers"
	"github.com/cafedex-backend/services"
	"github.com/clerk/clerk-sdk-go/v2"

)

type Application struct {
	Models services.Models
}

func main() {

	clerk.SetKey("sk_")

	mongoClient, err := db.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	services.New(mongoClient)

	log.Println("Server is running on port",8080)
	log.Fatal(http.ListenAndServe(":8080", handlers.CreateRouter()))
}
