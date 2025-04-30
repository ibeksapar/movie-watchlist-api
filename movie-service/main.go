package main

import (
	"log"
	"movie-service/handlers"
	"movie-service/middlewares"
	"movie-watchlist-api/db"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("DB_NAME", "watchlist")
	db.Connect()

	r := gin.Default()
	r.Use(middlewares.LoggerMiddleware())

	public := r.Group("/")
	{
		public.GET("/movies", handlers.GetMovies)
		public.GET("/movies/:id", handlers.GetMovieByID)
		public.GET("/genres", handlers.GetGenres)
		public.GET("/genres/:id", handlers.GetGenreByID)
	}

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/movies", handlers.CreateMovie)
		protected.PUT("/movies/:id", handlers.UpdateMovie)
		protected.DELETE("/movies/:id", handlers.DeleteMovie)
		protected.POST("/genres", handlers.CreateGenre)
		protected.PUT("/genres/:id", handlers.UpdateGenre)
		protected.DELETE("/genres/:id", handlers.DeleteGenre)
		protected.POST("/movies/:id/reviews", handlers.CreateReviewForMovie)
	}

	log.Println("Movie service running on http://localhost:8080")
	r.Run(":8080")
}
