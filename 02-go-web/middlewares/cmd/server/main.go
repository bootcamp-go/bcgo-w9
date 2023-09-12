package main

import (
	"app/cmd/server/handlers"
	"app/cmd/server/middlewares"
	"app/internal/movies/storage"
	"os"

	docs "app/documents/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Movies API
// @description This is a sample server for Movies API.
// @version 1
// @host localhost:8080
// @BasePath /api/v1

// @contact.name API Support
// @contact.email random@gmail.com
func main() {
	// env
	godotenv.Load()

	// dependencies
	db := []*storage.Movie{
		{
			Title: "Titanic",
			Year:  1997,
			Genre: "Drama",
		},
		{
			Title: "Crawl",
			Year:  2019,
			Genre: "Horror",
		},
	}
	st := storage.NewStorageMovieInMem(db)
	hd := handlers.NewMovieHandlers(st)

	// server
	rt := gin.New()
	rt.Use(
		// -> middlewares
		gin.Recovery(),
		middlewares.Logger(),
	)
	// -> routes / endpoints
	api := rt.Group("/api/v1")
	api.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "Pong")
	})
	// -> docs swaggo
	docs.SwaggerInfo.Host = os.Getenv("SERVER_ADDR")
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// -> movies
	mv := api.Group("/movies")
	mv.Use(
		// -> middlewares
		middlewares.Authenticator(),
	)
	{
		// -> get movies
		mv.GET("/", hd.GetMovies())
		// -> get movie by title
		mv.GET("/:title", hd.GetMovieByTitle())
	}
	
	if err := rt.Run(os.Getenv("SERVER_ADDR")); err != nil {
		panic(err)
	}	
}