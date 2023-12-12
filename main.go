package main

import (
	
	"gofr.dev/pkg/gofr"
	
	"github.com/rastogi26/gofr-netflix/routers"
)
func main()  {
	app := gofr.New()

	// Bypass header validation during API calls
	app.Server.ValidateHeaders = false

	// Setup Netflix routes
	routers.SetupNetflixRoutes(app)

	// Start the application
	app.Start()

}