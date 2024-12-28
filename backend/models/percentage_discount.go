package models

type PercentageDiscount struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Percentage float64 `json:"percentage"`
}
