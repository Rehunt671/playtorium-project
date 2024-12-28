package models

type FixedAmountDiscount struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Amount float64 `json:"amount"`
}
