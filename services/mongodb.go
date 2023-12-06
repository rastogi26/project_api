package services

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var (
	Client *mongo.Client
	DbName = "netflix"
)

func init() {

	// MongoDB connection setup
	connectionString := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println("MongoDB client created successfully")
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println("Ping to MongoDB successful")
	fmt.Println("Connected to MongoDB!")

	Client = client
	fmt.Printf("Client: %v\n", Client)
}

