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
	w.Header().Set("Content-Type", "application/json")

	var genres []models.Genre

	if err := db.DB.Preload("Movies").Find(&genres).Error; err != nil {
		http.Error(w, "Failed to fetch genres", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(genres)
}

func GetGenreByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var genre models.Genre

	if err := db.DB.Preload("Movies").First(&genre, id).Error; err != nil {
		http.Error(w, "Genre not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(genre)
}

func CreateGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

func UpdateGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var genre models.Genre

	if err := db.DB.First(&genre, id).Error; err != nil {
		http.Error(w, "Genre not found", http.StatusNotFound)
		return
	}

	var updatedGenre models.Genre
	if err := json.NewDecoder(r.Body).Decode(&updatedGenre); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if updatedGenre.Name != "" {
		genre.Name = updatedGenre.Name
	}

	if updatedGenre.Description != "" {
		genre.Description = updatedGenre.Description
	}


	if err := db.DB.Save(&genre).Error; err != nil {
		http.Error(w, "Failed to update genre", http.StatusInternalServerError)
		return
	}

	db.DB.Preload("Movies").First(&genre)

	json.NewEncoder(w).Encode(genre)
}

func DeleteGenre(w http.ResponseWriter, r *http.Request) {	
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := db.DB.Delete(&models.Genre{}, id).Error; err != nil {
		http.Error(w, "Failed to delete genre", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
