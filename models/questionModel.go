package models

type Question struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Value string `json:"value" gorm:"not null"`
}