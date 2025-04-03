package handlers

import (
	"encoding/json"
	"movie-watchlist-api/db"
	"movie-watchlist-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil || review.Score < 1 || review.Score > 10 || review.Content == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var movie models.Movie
	if err := db.DB.First(&movie, review.MovieID).Error; err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	db.DB.Create(&review)
	recalculateRating(review.MovieID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(review)
}

func GetReviews(w http.ResponseWriter, r *http.Request) {
	var reviews []models.Review
	db.DB.Find(&reviews)
	json.NewEncoder(w).Encode(reviews)
}

func GetReviewByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var review models.Review

	if err := db.DB.First(&review, id).Error; err != nil {
		http.Error(w, "Review not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(review)
}

func UpdateReview(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var review models.Review

	if err := db.DB.First(&review, id).Error; err != nil {
		http.Error(w, "Review not found", http.StatusNotFound)
		return
	}

	var updated models.Review
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil || updated.Score < 1 || updated.Score > 10 || updated.Content == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	review.Content = updated.Content
	review.Score = updated.Score
	db.DB.Save(&review)
	recalculateRating(review.MovieID)

	json.NewEncoder(w).Encode(review)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var review models.Review

	if err := db.DB.First(&review, id).Error; err != nil {
		http.Error(w, "Review not found", http.StatusNotFound)
		return
	}

	db.DB.Delete(&review)
	recalculateRating(review.MovieID)

	w.WriteHeader(http.StatusNoContent)
}

func recalculateRating(movieID uint) {
	var reviews []models.Review
	db.DB.Where("movie_id = ?", movieID).Find(&reviews)

	if len(reviews) == 0 {
		db.DB.Model(&models.Movie{}).Where("id = ?", movieID).Update("rating", 0)
		return
	}

	var total int
	for _, r := range reviews {
		total += r.Score
	}

	average := float64(total) / float64(len(reviews))
	db.DB.Model(&models.Movie{}).Where("id = ?", movieID).Update("rating", average)
}
