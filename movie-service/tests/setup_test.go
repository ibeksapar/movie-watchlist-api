package tests

import (
	"movie-service/movie_db"
	"movie-watchlist-api/db"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func TestMain(m *testing.M) {
	os.Setenv("DB_NAME", "watchlist_test")
	db.Connect()

	databaseURL := "postgres://postgres:postgres@localhost:2345/watchlist_test?sslmode=disable"
	mg, err := migrate.New(
		"file://../db/migrations",
		databaseURL,
	)

	if err != nil {
		panic("Failed to setup migration for test DB: " + err.Error())
	}

	if err := mg.Up(); err != nil && err != migrate.ErrNoChange {
		panic("Failed to apply migrations: " + err.Error())
	}

	movie_db.Seed()
	os.Exit(m.Run())
}
