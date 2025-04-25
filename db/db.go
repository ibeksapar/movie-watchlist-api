package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "watchlist" 
	}

	dsn := fmt.Sprintf("host=localhost user=postgres password=postgres dbname=%s port=2345 sslmode=disable", dbName)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to PostgreSQL:", dbName)
}
