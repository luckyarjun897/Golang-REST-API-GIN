package main

import (
	"api/db"
	"api/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovies(c *gin.Context) {
	fmt.Println("started GetMovies...")
	movies := db.GetMovies()
	c.IndentedJSON(http.StatusOK, movies)
}

func GetMovieByName(c *gin.Context) {
	fmt.Println("Started GetMovieByName....")
	name := c.Param("name")
	fmt.Printf("name=[%v]", name)
	present, movie := db.GetMovieByName(name)
	if !present {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie Not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, movie)
}

func AddMovie(c *gin.Context) {
	fmt.Println("Started AddMovie...")
	var newmovie types.Movies
	err := c.BindJSON(&newmovie)
	if err != nil {
		fmt.Errorf("Error while binding the new movie ,err = %v", err)
	}
	db.InsertMovie(newmovie)
	movies := db.GetMovies()
	c.IndentedJSON(http.StatusCreated, movies)
}

func UpdateMovie(c *gin.Context) {
	fmt.Println("Started UpdateMovie...")
	name := c.Param("name")
	var newmovie types.Movies
	err := c.BindJSON(&newmovie)
	if err != nil {
		fmt.Errorf("Error while binding the new movie ,err = %v", err)
	}
	present := db.UpdateMovie(name, newmovie)
	if !present {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Given movie is not in movies list"})
		return
	}
	movies := db.GetMovies()
	c.IndentedJSON(http.StatusOK, movies)
}

func DeleteMovie(c *gin.Context) {
	fmt.Println("started DeleteMovie...")
	name := c.Param("name")
	present := db.DeleteMovie(name)
	if !present {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Given movie is not in movies list"})
		return
	}
	movies := db.GetMovies()
	c.IndentedJSON(http.StatusOK, movies)
}

func main() {
	db.DBConnect()
	defer db.DB.Close()
	server := gin.Default()
	server.GET("/movies", GetMovies)
	server.GET("/movies/:name", GetMovieByName)
	server.POST("/movies", AddMovie)
	server.DELETE("/movies/:name", DeleteMovie)
	server.PUT("/movies/:name", UpdateMovie)
	fmt.Println("Starting localhost 8080")
	server.Run("localhost:8080")
}
