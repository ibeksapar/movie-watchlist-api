package handlers

import (
	"fmt"
	"movie-service/models"
	"movie-watchlist-api/db"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMovies(c *gin.Context) {
	var movies []models.Movie

	genreID := c.DefaultQuery("genre_id", "")
	title := c.DefaultQuery("title", "")
	minRating := c.DefaultQuery("min_rating", "")
	maxRating := c.DefaultQuery("max_rating", "")
	sortBy := c.DefaultQuery("sort", "")

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

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
	default:
		query = query.Order("id DESC")
	}

	 var total int64
		if err := query.Session(&gorm.Session{}).Model(&models.Movie{}).Count(&total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count movies"})
        return
    }

	if err := query.Limit(limit).Offset(offset).Find(&movies).Error; err != nil {
         c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
         return
     }

    c.JSON(http.StatusOK, gin.H{
        "movies": movies,
        "total":  total,
    })
}


func GetMovieByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var movie models.Movie
    result := db.DB.Preload("Genre").Preload("Reviews").First(&movie, id)

    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
        return
    }

    c.JSON(http.StatusOK, movie)
}

func CreateMovie(c *gin.Context) {
    var movie models.Movie

    if err := c.ShouldBindJSON(&movie); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    if movie.Title == "" || movie.GenreID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Title and GenreID are required"})
        return
    }

    if movie.Rating < 1 || movie.Rating > 10 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 10"})
        return
    }

    if err := db.DB.Create(&movie).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving movie"})
        return
    }

    c.JSON(http.StatusCreated, movie)
}

func UpdateMovie(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var movie models.Movie

    if err := db.DB.First(&movie, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
        return
    }

    var updatedMovie models.Movie
    if err := c.ShouldBindJSON(&updatedMovie); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    movie.Title = updatedMovie.Title
    movie.Rating = updatedMovie.Rating
    movie.GenreID = updatedMovie.GenreID

    if err := db.DB.Save(&movie).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
        return
    }

    c.JSON(http.StatusOK, movie)
}

func DeleteMovie(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    db.DB.Delete(&models.Movie{}, id)
    c.Status(http.StatusNoContent)
}

func UploadCover(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var movie models.Movie
    if err := db.DB.First(&movie, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
        return
    }

    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
        return
    }

    filename := fmt.Sprintf("%d_%s", movie.ID, filepath.Base(file.Filename))
    savePath := filepath.Join("uploads", filename)
    if err := c.SaveUploadedFile(file, savePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }

    url := "/uploads/" + filename

    movie.CoverURL = url
    if err := db.DB.Save(&movie).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"url": url})
}

func UploadGeneral(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
        return
    }

    filename := filepath.Base(file.Filename)
    savePath := filepath.Join("uploads", filename)
    if err := c.SaveUploadedFile(file, savePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }

    url := "/uploads/" + filename
    c.JSON(http.StatusOK, gin.H{"url": url})
}