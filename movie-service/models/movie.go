package models

type Movie struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	Title    string   `json:"title"`
	GenreID  uint     `json:"genre_id"`
	Genre    Genre    `json:"genre"`
	Rating   float64  `json:"rating"`
	Reviews  []Review `json:"reviews"`
	CoverURL string   `json:"cover_url"`
}
