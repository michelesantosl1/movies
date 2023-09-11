package main

import (
	"movies/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/movie", handlers.CreateMovie)
	r.GET("/movies", handlers.GetMovies)
	r.GET("/movie/:id", handlers.GetMoviesByID)
	r.PUT("/movie/:id", handlers.UpdateMovie)
	r.DELETE("/movie/:id", handlers.DeleteMovie)

	r.Run(":8080")
}
