package main

import (
	"log"
	"movie-watchlist-api/db"
	"os"
	"user-service/handlers"
	"user-service/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Setenv("DB_NAME", "watchlist")
	db.Connect()

	r := gin.Default()
	r.Use(middlewares.LoggerMiddleware())

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/validate", handlers.ValidateToken)

	log.Println("User service running on http://localhost:8081")
	r.Run(":8081")
}
