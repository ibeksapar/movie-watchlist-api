package handlers

import (
	"movie-service/models"
	"movie-watchlist-api/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil || review.Score < 1 || review.Score > 10 || review.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var movie models.Movie
	if err := db.DB.First(&movie, review.MovieID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if err := db.DB.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}

	recalculateRating(review.MovieID)

	c.JSON(http.StatusCreated, review)
}

func CreateReviewForMovie(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var movie models.Movie
	if err := db.DB.First(&movie, movieID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if review.Score < 1 || review.Score > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Score must be between 1 and 10"})
		return
	}

	review.MovieID = uint(movieID)

	if err := db.DB.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create review"})
		return
	}

	recalculateRating(review.MovieID)

	c.JSON(http.StatusCreated, review)
}

func GetReviewsByMovie(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var reviews []models.Review
	if err := db.DB.Where("movie_id = ?", movieID).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func GetReviews(c *gin.Context) {
	var reviews []models.Review
	if err := db.DB.Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func GetReviewByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var review models.Review

	if err := db.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	c.JSON(http.StatusOK, review)
}

func UpdateReview(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var review models.Review

	if err := db.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	var updated models.Review
	if err := c.ShouldBindJSON(&updated); err != nil || updated.Score < 1 || updated.Score > 10 || updated.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	review.Content = updated.Content
	review.Score = updated.Score
	if err := db.DB.Save(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update review"})
		return
	}

	recalculateRating(review.MovieID)

	c.JSON(http.StatusOK, review)
}

func DeleteReview(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var review models.Review

	if err := db.DB.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	if err := db.DB.Delete(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete review"})
		return
	}

	recalculateRating(review.MovieID)

	c.Status(http.StatusNoContent)
}

func recalculateRating(movieID uint) {
	var reviews []models.Review
	if err := db.DB.Where("movie_id = ?", movieID).Find(&reviews).Error; err != nil {
		return
	}

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
