package handlers

import (
	"movie-service/models"
	"net/http"
	"strconv"

	"movie-watchlist-api/db"

	"github.com/gin-gonic/gin"
)

func GetGenres(c *gin.Context) {
    var genres []models.Genre

    if err := db.DB.Preload("Movies").Find(&genres).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch genres"})
        return
    }

    c.JSON(http.StatusOK, genres)
}

func GetGenreByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var genre models.Genre

    if err := db.DB.Preload("Movies").First(&genre, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
        return
    }

    c.JSON(http.StatusOK, genre)
}

func CreateGenre(c *gin.Context) {
    var genre models.Genre

    if err := c.ShouldBindJSON(&genre); err != nil || genre.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := db.DB.Create(&genre).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create genre"})
        return
    }

    c.JSON(http.StatusCreated, genre)
}

func UpdateGenre(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var genre models.Genre

    if err := db.DB.First(&genre, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
        return
    }

    var updatedGenre models.Genre
    if err := c.ShouldBindJSON(&updatedGenre); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if updatedGenre.Name != "" {
        genre.Name = updatedGenre.Name
    }

    if updatedGenre.Description != "" {
        genre.Description = updatedGenre.Description
    }

    if err := db.DB.Save(&genre).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update genre"})
        return
    }

    c.JSON(http.StatusOK, genre)
}

func DeleteGenre(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))

    if err := db.DB.Delete(&models.Genre{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete genre"})
        return
    }

    c.Status(http.StatusNoContent)
}
