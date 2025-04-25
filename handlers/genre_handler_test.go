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

func TestCreateGenre(t *testing.T) {
	newGenre := models.Genre{Name: "Horror", Description: "Scary movies"}
	body, _ := json.Marshal(newGenre)

	req, _ := http.NewRequest("POST", "/genres", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateGenre)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var result models.Genre
	_ = json.NewDecoder(rr.Body).Decode(&result)
	assert.Equal(t, newGenre.Name, result.Name)
}

func TestGetGenres(t *testing.T) {
	req, _ := http.NewRequest("GET", "/genres", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetGenres)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var genres []models.Genre
	_ = json.NewDecoder(rr.Body).Decode(&genres)
	assert.NotEmpty(t, genres)
}

func TestGetGenreByID(t *testing.T) {
	genre := CreateTestGenre()

	req := httptest.NewRequest("GET", fmt.Sprintf("/genres/%d", genre.ID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprint(genre.ID)})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetGenreByID)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result models.Genre
	_ = json.NewDecoder(rr.Body).Decode(&result)
	assert.Equal(t, genre.ID, result.ID)
}

func TestUpdateGenre(t *testing.T) {
	genre := CreateTestGenre()
	updated := models.Genre{Name: "Updated", Description: "Changed"}
	body, _ := json.Marshal(updated)

	req := httptest.NewRequest("PUT", fmt.Sprintf("/genres/%d", genre.ID), bytes.NewBuffer(body))
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprint(genre.ID)})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateGenre)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result models.Genre
	_ = json.NewDecoder(rr.Body).Decode(&result)
	assert.Equal(t, updated.Name, result.Name)
}

func TestDeleteGenre(t *testing.T) {
	genre := CreateTestGenre()

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/genres/%d", genre.ID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprint(genre.ID)})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteGenre)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
