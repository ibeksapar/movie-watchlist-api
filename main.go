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
	db.DB.AutoMigrate(&models.Genre{}, &models.Movie{})

	r := mux.NewRouter()

	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMovieByID).Methods("GET")
	r.HandleFunc("/movies", handlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")

	r.HandleFunc("/genres", handlers.GetGenres).Methods("GET")
	r.HandleFunc("/genres/{id}", handlers.GetGenreByID).Methods("GET")
	r.HandleFunc("/genres", handlers.CreateGenre).Methods("POST")


	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
