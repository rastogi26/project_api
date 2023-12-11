package services

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)


var (
	Client *mongo.Client
	DbName string
)

func init() {

	// MongoDB connection setup
	// connectionString := "mongodb://localhost:27017"
	err := godotenv.Load("config/.env") 
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DbName = os.Getenv("MONGO_DB_NAME")

	connectionString := os.Getenv("MONGO_DB_CONNECTION_STRING")
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

