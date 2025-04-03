package handlers

import (
	"encoding/json"
	"movie-watchlist-api/db"
	"movie-watchlist-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetGenres(w http.ResponseWriter, r *http.Request) {
	var genres []models.Genre

	if err := db.DB.Preload("Movies").Find(&genres).Error; err != nil {
		http.Error(w, "Failed to fetch genres", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(genres)
}

func GetGenreByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var genre models.Genre

	if err := db.DB.Preload("Movies").First(&genre, id).Error; err != nil {
		http.Error(w, "Genre not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(genre)
}



func CreateGenre(w http.ResponseWriter, r *http.Request) {
	var genre models.Genre

	if err := json.NewDecoder(r.Body).Decode(&genre); err != nil || genre.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := db.DB.Create(&genre).Error; err != nil {
		http.Error(w, "Failed to create genre", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(genre)
}
