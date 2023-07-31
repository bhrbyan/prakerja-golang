package handler

import (
	"movie-api/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateMovieHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		movie := new(model.Movie)
		if err := c.Bind(movie); err != nil {
			return err
		}
		if err := model.CreateMovie(db, movie); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, movie)
	}
}

func GetMoviesHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var movies []model.Movie
		if err := model.GetMovies(db, &movies); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, movies)
	}
}

func GetMovieHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		movie := new(model.Movie)
		id := c.Param("id")
		if err := model.GetMovie(db, movie, id); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, movie)
	}
}

func UpdateMovieHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		movie := new(model.Movie)
		if err := c.Bind(movie); err != nil {
			return err
		}
		if err := model.UpdateMovie(db, movie); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, movie)
	}
}

func DeleteMovieHandler(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if err := model.DeleteMovie(db, id); err != nil {
			return err
		}
		return c.NoContent(http.StatusNoContent)
	}
}
