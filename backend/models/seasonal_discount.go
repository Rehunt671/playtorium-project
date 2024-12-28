package models

type SeasonalDiscount struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	EveryXTHB    float64 `json:"every_xthb"`
	DiscountYTHB float64 `json:"discount_ythb"`
}
