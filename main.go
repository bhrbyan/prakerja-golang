package main

import (
	"movie-api/handler"
	"movie-api/model"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Initialize SQLite database with GORM
	db, err := gorm.Open(sqlite.Open("movies.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Movie{})

	// Create Movie Routes
	e.POST("/movies", handler.CreateMovieHandler(db))
	e.GET("/movies", handler.GetMoviesHandler(db))
	e.GET("/movies/:id", handler.GetMovieHandler(db))
	e.PUT("/movies/:id", handler.UpdateMovieHandler(db))
	e.DELETE("/movies/:id", handler.DeleteMovieHandler(db))

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
