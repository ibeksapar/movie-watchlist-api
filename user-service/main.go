package main

import (
	"log"
	"movie-watchlist-api/db"
	"os"
	"user-service/handlers"
	"user-service/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("DB_NAME", "watchlist")
	db.Connect()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.Use(middlewares.LoggerMiddleware())

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/validate", handlers.ValidateToken)

	log.Println("User service running on http://localhost:8081")
	r.Run(":8081")
}
