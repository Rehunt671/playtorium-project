package models

import "gorm.io/gorm"

type PointDiscount struct {
	gorm.Model
	Point float64 `json:"point"` // 1 point = 1 THB
}
