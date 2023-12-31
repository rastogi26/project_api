package handlers

import (
	

	"github.com/rastogi26/gofr-netflix/models"
	"github.com/rastogi26/gofr-netflix/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gofr.dev/pkg/gofr"
)

type NetflixHandler struct{}

func NewNetflixHandler() NetflixHandler {
	return NetflixHandler{}
}

func (h NetflixHandler) GetMyAllMovies(ctx *gofr.Context) (interface{}, error) {
    movies := services.GetAllMovies()
	return movies, nil
}


func (h NetflixHandler) AddMovie(ctx *gofr.Context) (interface{}, error) {
	var movie models.Netflix
	err := ctx.Bind(&movie)
	if err != nil {
		return nil, err
	}
    services.InsertOneMovie(movie)
	return "Movie added successfully", nil
}

func (h NetflixHandler) WatchedMovie(ctx *gofr.Context) (interface{}, error) {
	// log.Println("Call1 send")
	movieId := ctx.PathParam("id")
	services.UpdateOneMovie(movieId)
	return "Movie marked as watched", nil
}

func (h NetflixHandler) RemoveAllMovies(ctx *gofr.Context) (interface{}, error) {
	deleteCount := services.DeleteAllMovie()
	return deleteCount, nil
}


func (h NetflixHandler) RemoveMovie(ctx *gofr.Context) (interface{}, error) {
	// log.Println("Call2 send")
	movieId := ctx.PathParam("id")

	// Check if movieId is a valid ObjectId
    _, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		return nil,err
	}
      services.DeleteOneMovie(movieId)
	  return "Movie removed successfully",nil
}
