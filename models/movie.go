package models

type Movie struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	Title   string  `json:"title"`
	GenreID uint    `json:"genre_id"`
	Rating  float64 `json:"rating"`
}