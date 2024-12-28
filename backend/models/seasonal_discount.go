package models

import "gorm.io/gorm"

type SeasonalDiscount struct {
	gorm.Model
	EveryXTHB    float64 `json:"every_xthb"`
	DiscountYTHB float64 `json:"discount_ythb"`
}
