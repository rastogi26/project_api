package routers

import (
	

	"github.com/rastogi26/gofr-netflix/handlers"
	"gofr.dev/pkg/gofr"
)

func SetupNetflixRoutes(app *gofr.Gofr) {
	netflixHandler := handlers.NewNetflixHandler()
    
	app.GET("/mymovies", netflixHandler.GetMyAllMovies)
	app.POST("/addmovie", netflixHandler.AddMovie)
	app.PUT("/watched/{id}", netflixHandler.WatchedMovie)
	app.DELETE("/removeallmovies", netflixHandler.RemoveAllMovies)
	app.DELETE("/removemovie/{id}", netflixHandler.RemoveMovie)
}



