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
	w.Header().Set("Content-Type", "application/json")

	var movies []models.Movie

	genreID := r.URL.Query().Get("genre_id")
	title := r.URL.Query().Get("title")
	minRating := r.URL.Query().Get("min_rating")
	maxRating := r.URL.Query().Get("max_rating")
	sortBy := r.URL.Query().Get("sort")

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	query := db.DB.Preload("Genre").Preload("Reviews")

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

	if err := query.Limit(limit).Offset(offset).Find(&movies).Error; err != nil {
		http.Error(w, "Failed to fetch movies", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movies)
}

func GetMovieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var movie models.Movie
	result := db.DB.Preload("Genre").Preload("Reviews").First(&movie, id)

	if result.Error != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(movie)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if movie.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	if movie.Rating < 1 || movie.Rating > 10 {
		http.Error(w, "Rating must be between 1 and 10", http.StatusBadRequest)
		return
	}

	var genre models.Genre
	if err := db.DB.First(&genre, movie.GenreID).Error; err != nil {
		http.Error(w, "Genre does not exist", http.StatusBadRequest)
		return
	}

	db.DB.Create(&movie)
	w.WriteHeader(http.StatusCreated) 
	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var movie models.Movie

	if err := db.DB.First(&movie, id).Error; err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	var updatedMovie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if updatedMovie.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	if updatedMovie.Rating < 1 || updatedMovie.Rating > 10 {
		http.Error(w, "Rating must be between 1 and 10", http.StatusBadRequest)
		return
	}

	var genre models.Genre
	if err := db.DB.First(&genre, updatedMovie.GenreID).Error; err != nil {
		http.Error(w, "Genre does not exist", http.StatusBadRequest)
		return
	}

	movie.Title = updatedMovie.Title
	movie.Rating = updatedMovie.Rating
	movie.GenreID = updatedMovie.GenreID

	if err := db.DB.Save(&movie).Error; err != nil {
		http.Error(w, "Failed to update movie", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movie)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	db.DB.Delete(&models.Movie{}, id)
	w.WriteHeader(http.StatusNoContent)
}
