package models

type Genre struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Name   string  `json:"name"`
	Movies []Movie `json:"movies"`
}
