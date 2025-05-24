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
    host := os.Getenv("DB_HOST")
    if host == "" {
        host = "localhost"
    }
    
    port := os.Getenv("DB_PORT")
    if port == "" {
        port = "2345"
    }

    user := os.Getenv("DB_USER")
    if user == "" {
        user = "postgres"
    }

    password := os.Getenv("DB_PASSWORD")
    if password == "" {
        password = "postgres"
    }

    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "watchlist"
    }

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        host, user, password, dbName, port,
    )

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    log.Println("Connected to PostgreSQL:", host, port, dbName)
}
