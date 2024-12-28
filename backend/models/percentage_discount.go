package models

import "gorm.io/gorm"

type PercentageDiscount struct {
	gorm.Model
	Percentage float64 `json:"percentage"`
}
