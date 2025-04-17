package main

import (
	"log"
	"movie-watchlist-api/db"
	"movie-watchlist-api/handlers"
	"movie-watchlist-api/middlewares"
	"movie-watchlist-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db.Connect()
	db.DB.AutoMigrate(&models.Genre{}, &models.Movie{}, &models.Review{}, &models.User{})

	db.Seed() 

	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	protected := r.PathPrefix("/protected").Subrouter()
	protected.Use(middlewares.Authenticate)

	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMovieByID).Methods("GET")
	protected.HandleFunc("/movies", handlers.CreateMovie).Methods("POST")
	protected.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	protected.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")
	protected.HandleFunc("/movies/{id}/reviews", handlers.GetReviewsByMovie).Methods("GET")
	protected.HandleFunc("/movies/{id}/reviews", handlers.CreateReviewForMovie).Methods("POST")

	r.HandleFunc("/genres", handlers.GetGenres).Methods("GET")
	r.HandleFunc("/genres/{id}", handlers.GetGenreByID).Methods("GET")
	protected.HandleFunc("/genres", handlers.CreateGenre).Methods("POST")
	protected.HandleFunc("/genres/{id}", handlers.UpdateGenre).Methods("PUT")
	protected.HandleFunc("/genres/{id}", handlers.DeleteGenre).Methods("DELETE")

	protected.HandleFunc("/reviews", handlers.CreateReview).Methods("POST")
	r.HandleFunc("/reviews", handlers.GetReviews).Methods("GET")
	r.HandleFunc("/reviews/{id}", handlers.GetReviewByID).Methods("GET")
	protected.HandleFunc("/reviews/{id}", handlers.UpdateReview).Methods("PUT")
	protected.HandleFunc("/reviews/{id}", handlers.DeleteReview).Methods("DELETE")


	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
