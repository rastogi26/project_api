package services

import (
	"context"
	"fmt"
	"log"

	"github.com/rastogi26/gofr-netflix/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



var collection *mongo.Collection


func init() {
	// MongoDB collection setup
	collection = Client.Database(DbName).Collection("watchlist")
}


func InsertOneMovie(movie models.Netflix) {
	_, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		// Handle error
		log.Fatal(err)
	}
}

func GetAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		// Handle error
		log.Fatal(err)
	}

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			// Handle error
			log.Fatal(err)
			
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

//  this function updates the watched status from false to true.
func UpdateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
     
	//  log.Println("call1 received")
	//  log.Printf("ID : %v\n", id)
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		// Handle error
		log.Fatal(err)
	}
	
}

func DeleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}

	 
     //  log.Println("call2 received")
	//  log.Printf("ID : %v\n", id)
	 deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		// Handle error
		log.Fatal(err)
	}
	fmt.Println("Movie got delete with delete count: ", deleteCount)

}

func DeleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		// Handle error
		log.Fatal(err)
	}

	return deleteResult.DeletedCount
}




