package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"movie-watchlist-api/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateMovie(t *testing.T) {
	genre := CreateTestGenre()
	newMovie := models.Movie{Title: "New Movie", GenreID: genre.ID, Rating: 9}
	body, _ := json.Marshal(newMovie)

	req, _ := http.NewRequest("POST", "/movies", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateMovie)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var result models.Movie
	_ = json.NewDecoder(rr.Body).Decode(&result)
	assert.Equal(t, newMovie.Title, result.Title)
}

func TestGetMovies(t *testing.T) {
	req, _ := http.NewRequest("GET", "/movies", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMovies)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var movies []models.Movie
	_ = json.NewDecoder(rr.Body).Decode(&movies)
	assert.NotEmpty(t, movies)
}

func TestGetMovieByID(t *testing.T) {
	genre := CreateTestGenre()
	movie := CreateTestMovie(genre.ID)

	req := httptest.NewRequest("GET", fmt.Sprintf("/movies/%d", movie.ID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprint(movie.ID)})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMovieByID)
	handler.ServeHTTP(rr, req)

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

	req := httptest.NewRequest("PUT", fmt.Sprintf("/movies/%d", movie.ID), bytes.NewBuffer(body))
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprint(movie.ID)})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateMovie)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result models.Movie
	_ = json.NewDecoder(rr.Body).Decode(&result)
	assert.Equal(t, updated.Title, result.Title)
}

func TestDeleteMovie(t *testing.T) {
	genre := CreateTestGenre()
	movie := CreateTestMovie(genre.ID)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/movies/%d", movie.ID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprint(movie.ID)})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteMovie)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
