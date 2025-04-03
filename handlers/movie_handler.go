package handlers

import (
	"encoding/json"
	"movie-watchlist-api/db"
	"movie-watchlist-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie
	db.DB.Find(&movies)
	json.NewEncoder(w).Encode(movies)
}

func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var movie models.Movie
	result := db.DB.First(&movie, id)
	if result.Error != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(movie)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)

	if movie.Title == "" || movie.Rating <= 0 {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	db.DB.Create(&movie)
	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var movie models.Movie
	if err := db.DB.First(&movie, id).Error; err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	json.NewDecoder(r.Body).Decode(&movie)
	db.DB.Save(&movie)
	json.NewEncoder(w).Encode(movie)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	db.DB.Delete(&models.Movie{}, id)
	w.WriteHeader(http.StatusNoContent)
}
