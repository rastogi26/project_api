package services

import (
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

}