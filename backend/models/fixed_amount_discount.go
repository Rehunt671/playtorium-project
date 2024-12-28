package models

import "gorm.io/gorm"

type FixedAmountDiscount struct {
	gorm.Model
	Amount float64 `json:"amount"`
}
