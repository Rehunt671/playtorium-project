package models

type PointDiscount struct {
	ID    uint    `json:"id" gorm:"primaryKey"`
	Point float64 `json:"point"` // 1 point = 1 THB
}
