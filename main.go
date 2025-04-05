package main

import (
	"log"
	"movie-watchlist-api/db"
	"movie-watchlist-api/handlers"
	"movie-watchlist-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect()
	db.DB.AutoMigrate(&models.Genre{}, &models.Movie{}, &models.Review{})

	db.Seed() 

	r := mux.NewRouter()

	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMovieByID).Methods("GET")
	r.HandleFunc("/movies", handlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}/reviews", handlers.GetReviewsByMovie).Methods("GET")
	r.HandleFunc("/movies/{id}/reviews", handlers.CreateReviewForMovie).Methods("POST")

	r.HandleFunc("/genres", handlers.GetGenres).Methods("GET")
	r.HandleFunc("/genres/{id}", handlers.GetGenreByID).Methods("GET")
	r.HandleFunc("/genres", handlers.CreateGenre).Methods("POST")
	r.HandleFunc("/genres/{id}", handlers.UpdateGenre).Methods("PUT")
	r.HandleFunc("/genres/{id}", handlers.DeleteGenre).Methods("DELETE")

	r.HandleFunc("/reviews", handlers.CreateReview).Methods("POST")
	r.HandleFunc("/reviews", handlers.GetReviews).Methods("GET")
	r.HandleFunc("/reviews/{id}", handlers.GetReviewByID).Methods("GET")
	r.HandleFunc("/reviews/{id}", handlers.UpdateReview).Methods("PUT")
	r.HandleFunc("/reviews/{id}", handlers.DeleteReview).Methods("DELETE")


	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
