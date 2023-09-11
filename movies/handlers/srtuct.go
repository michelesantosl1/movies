package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// sctruct for movies
type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

var movies = []Movie{}

func CreateMovie(c *gin.Context) {
	var newMovie Movie
	if err := c.ShouldBindJSON(&newMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMovie.ID = len(movies) + 1
	movies = append(movies, newMovie)
	c.JSON(http.StatusCreated, newMovie)
}

func GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, movies)
}

func GetMoviesByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	for _, movie := range movies {
		if movie.ID == id {
			c.JSON(http.StatusOK, movie)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

func UpdateMovie(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var updatedMovie Movie
	if err := c.ShouldBindJSON(&updatedMovie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, movie := range movies {
		if movie.ID == id {
			movies[i] = updatedMovie
			c.JSON(http.StatusOK, updatedMovie)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

func DeleteMovie(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	for i, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}