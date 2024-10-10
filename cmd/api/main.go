package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cafedex-backend/handlers"
	"github.com/cafedex-backend/services"
	// "github.com/clerk/clerk-sdk-go/v2"
)

type Application struct {
	Models services.Models
}

func main() {

	// clerk.SetKey("sk_")

	// ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer cancel()

	// defer func() {
	// 	if err = mongoClient.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// services.New(mongoClient)

	fmt.Println("Server is running on port", 8080)
	log.Fatal(http.ListenAndServe(":8080", handlers.CreateRouter()))
}
