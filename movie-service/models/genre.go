package models

type Genre struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Movies      []Movie `json:"movies"`
}
