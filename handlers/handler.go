package handlers

import(
	"github.com/rastogi26/gofr-netflix/models"
	"github.com/rastogi26/gofr-netflix/services"
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