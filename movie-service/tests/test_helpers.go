package tests

import (
	"log"
	"movie-service/models"
	"movie-watchlist-api/db"
)

func CreateTestGenre() models.Genre {
	genre := models.Genre{Name: "Test Genre", Description: "For testing"}
	if err := db.DB.Create(&genre).Error; err != nil {
		log.Fatalf("Failed to create test genre: %v", err)
	}

	return genre
}

func CreateTestMovie(genreID uint) models.Movie {
	movie := models.Movie{Title: "Test Movie", GenreID: genreID, Rating: 7.5}
	if err := db.DB.Create(&movie).Error; err != nil {
		log.Fatalf("Failed to create test movie: %v", err)
	}
	
	return movie
}
