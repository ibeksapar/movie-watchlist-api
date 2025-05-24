package movie_db

import (
	"log"
	"movie-service/models"

	"movie-watchlist-api/db"
)

func Seed() {
	log.Println("Seeding database...")

	db.DB.Exec("DELETE FROM reviews")
	db.DB.Exec("DELETE FROM movies")
	db.DB.Exec("DELETE FROM genres")
	db.DB.Exec("ALTER SEQUENCE reviews_id_seq RESTART WITH 1")
	db.DB.Exec("ALTER SEQUENCE movies_id_seq RESTART WITH 1")
	db.DB.Exec("ALTER SEQUENCE genres_id_seq RESTART WITH 1")

	genres := []models.Genre{
		{Name: "Thriller", Description: "Suspenseful stories that keep you on edge"},
		{Name: "Science Fiction", Description: "Futuristic and scientific settings"},
		{Name: "Comedy", Description: "Light-hearted and humorous stories"},
		{Name: "Drama", Description: "Serious narratives that explore emotions"},
		{Name: "Action", Description: "Fast-paced plots with physical stunts"},
	}

	for _, g := range genres {
		db.DB.Create(&g)
	}

	movies := []models.Movie{
		{Title: "Inception", GenreID: 2, Rating: 0},
		{Title: "The Martian", GenreID: 2, Rating: 0},
		{Title: "The Hangover", GenreID: 3, Rating: 0},
		{Title: "Joker", GenreID: 4, Rating: 0},
		{Title: "John Wick", GenreID: 5, Rating: 0},
		{Title: "Gone Girl", GenreID: 1, Rating: 0},
		{Title: "Interstellar", GenreID: 2, Rating: 0},
		{Title: "The Office", GenreID: 3, Rating: 0},
		{Title: "The Room", GenreID: 4, Rating: 0},
		{Title: "Sharknado", GenreID: 5, Rating: 0},
	}

	for _, m := range movies {
		db.DB.Create(&m)
	}

	reviews := []models.Review{
		{MovieID: 1, Content: "Genius-level storytelling.", Score: 10},
		{MovieID: 1, Content: "Hard to follow but cool", Score: 7},

		{MovieID: 2, Content: "Scientific and inspiring!", Score: 8},
		{MovieID: 2, Content: "Slow in parts", Score: 6},

		{MovieID: 3, Content: "Super funny", Score: 9},

		{MovieID: 4, Content: "Disturbing yet powerful", Score: 10},
		{MovieID: 4, Content: "Too dark for me", Score: 5},
		{MovieID: 4, Content: "Brilliant performance!", Score: 9},

		{MovieID: 5, Content: "Non-stop adrenaline", Score: 8},

		{MovieID: 6, Content: "Kept me guessing", Score: 7},

		{MovieID: 7, Content: "Visually stunning", Score: 10},
		{MovieID: 7, Content: "Overrated", Score: 5},

		{MovieID: 8, Content: "Classic comedy!", Score: 8},
		{MovieID: 8, Content: "Outdated", Score: 4},

		{MovieID: 9, Content: "So bad, it's funny", Score: 2},
		{MovieID: 9, Content: "The worst acting ever", Score: 1},

		{MovieID: 10, Content: "Trashy fun", Score: 4},
	}

	for _, r := range reviews {
		db.DB.Create(&r)
	}

	var allMovies []models.Movie
	db.DB.Find(&allMovies)
	
	for _, movie := range allMovies {
		recalculateRating(movie.ID)
	}

	log.Println("Database seeded successfully.")
}

func recalculateRating(movieID uint) {
	var reviews []models.Review
	db.DB.Where("movie_id = ?", movieID).Find(&reviews)

	if len(reviews) == 0 {
		db.DB.Model(&models.Movie{}).Where("id = ?", movieID).Update("rating", 0)
		return
	}

	total := 0
	for _, r := range reviews {
		total += r.Score
	}

	avg := float64(total) / float64(len(reviews))
	db.DB.Model(&models.Movie{}).Where("id = ?", movieID).Update("rating", avg)
}
