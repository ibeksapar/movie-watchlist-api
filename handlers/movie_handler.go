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

	genreID := r.URL.Query().Get("genre_id")
	title := r.URL.Query().Get("title")
	minRating := r.URL.Query().Get("min_rating")
	maxRating := r.URL.Query().Get("max_rating")
	sortBy := r.URL.Query().Get("sort")

	query := db.DB.Preload("Genre")

	if genreID != "" {
		query = query.Where("genre_id = ?", genreID)
	}
	
	if title != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+title+"%")
	}

	if minRating != "" {
		query = query.Where("rating >= ?", minRating)
	}

	if maxRating != "" {
		query = query.Where("rating <= ?", maxRating)
	}

	switch sortBy {
		case "rating_asc":
			query = query.Order("rating ASC")
		case "rating_desc":
			query = query.Order("rating DESC")
		case "title_asc":
			query = query.Order("title ASC")
		case "title_desc":
			query = query.Order("title DESC")
	}

	if err := query.Find(&movies).Error; err != nil {
		http.Error(w, "Failed to fetch movies", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movies)
}


func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var movie models.Movie
	result := db.DB.Preload("Genre").First(&movie, id)

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
