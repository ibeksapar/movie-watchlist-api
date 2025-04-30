package models

type Review struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	MovieID uint   `json:"movie_id"`
	Content string `json:"content"`
	Score   int    `json:"score"`
}