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

func TestCreateGenre(t *testing.T) {
	r := gin.Default()
	r.POST("/genres", handlers.CreateGenre)

	newGenre := models.Genre{Name: "Horror", Description: "Scary movies"}
	body, _ := json.Marshal(newGenre)

	req, _ := http.NewRequest("POST", "/genres", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var result models.Genre
	_ = json.NewDecoder(rr.Body).Decode(&result)
	assert.Equal(t, newGenre.Name, result.Name)
}

func TestGetGenres(t *testing.T) {
	r := gin.Default()
	r.GET("/genres", handlers.GetGenres)

	req, _ := http.NewRequest("GET", "/genres", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var genres []models.Genre
	_ = json.NewDecoder(rr.Body).Decode(&genres)
	assert.NotEmpty(t, genres)
}

func TestGetGenreByID(t *testing.T) {
	genre := CreateTestGenre()

	r := gin.Default()
	r.GET("/genres/:id", handlers.GetGenreByID)

	req := httptest.NewRequest("GET", fmt.Sprintf("/genres/%d", genre.ID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprint(genre.ID)})

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result models.Genre
	_ = json.NewDecoder(rr.Body).Decode(&result)
	assert.Equal(t, genre.ID, result.ID)
}

func TestUpdateGenre(t *testing.T) {
	genre := CreateTestGenre()
	updated := models.Genre{Name: "Updated", Description: "Changed"}
	body, _ := json.Marshal(updated)

	r := gin.Default()
	r.PUT("/genres/:id", handlers.UpdateGenre)

	req := httptest.NewRequest("PUT", fmt.Sprintf("/genres/%d", genre.ID), bytes.NewBuffer(body))
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprint(genre.ID)})

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result models.Genre
	_ = json.NewDecoder(rr.Body).Decode(&result)
	assert.Equal(t, updated.Name, result.Name)
}

func TestDeleteGenre(t *testing.T) {
	genre := CreateTestGenre()

	r := gin.Default()
	r.DELETE("/genres/:id", handlers.DeleteGenre)

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/genres/%d", genre.ID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": fmt.Sprint(genre.ID)})

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
}
