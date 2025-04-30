package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"movie-service/handlers"
	"movie-service/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateMovie(t *testing.T) {
	genre := CreateTestGenre()
	newMovie := models.Movie{Title: "New Movie", GenreID: genre.ID, Rating: 9}
	body, _ := json.Marshal(newMovie)

	r := gin.Default()
	r.POST("/movies", handlers.CreateMovie)

	req, _ := http.NewRequest("POST", "/movies", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var result models.Movie
	_ = json.NewDecoder(rr.Body).Decode(&result)
	assert.Equal(t, newMovie.Title, result.Title)
}

func TestGetMovies(t *testing.T) {
	r := gin.Default()
	r.GET("/movies", handlers.GetMovies)

	req, _ := http.NewRequest("GET", "/movies", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var movies []models.Movie
	_ = json.NewDecoder(rr.Body).Decode(&movies)
	assert.NotEmpty(t, movies)
}

func TestGetMovieByID(t *testing.T) {
	genre := CreateTestGenre()
	movie := CreateTestMovie(genre.ID)

	r := gin.Default()
	r.GET("/movies/:id", handlers.GetMovieByID)

	req := httptest.NewRequest("GET", fmt.Sprintf("/movies/%d", movie.ID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprint(movie.ID)})

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result models.Movie
	_ = json.NewDecoder(rr.Body).Decode(&result)
	assert.Equal(t, movie.ID, result.ID)
}

func TestUpdateMovie(t *testing.T) {
	genre := CreateTestGenre()
	movie := CreateTestMovie(genre.ID)
	updated := models.Movie{Title: "Updated Title", GenreID: genre.ID, Rating: 8}
	body, _ := json.Marshal(updated)

	r := gin.Default()
	r.PUT("/movies/:id", handlers.UpdateMovie)

	req := httptest.NewRequest("PUT", fmt.Sprintf("/movies/%d", movie.ID), bytes.NewBuffer(body))
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprint(movie.ID)})

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result models.Movie
	_ = json.NewDecoder(rr.Body).Decode(&result)
	assert.Equal(t, updated.Title, result.Title)
}

func TestDeleteMovie(t *testing.T) {
	genre := CreateTestGenre()
	movie := CreateTestMovie(genre.ID)

	r := gin.Default()
	r.DELETE("/movies/:id", handlers.DeleteMovie)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/movies/%d", movie.ID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprint(movie.ID)})

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
