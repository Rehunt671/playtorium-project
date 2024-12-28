package models

import "gorm.io/gorm"

type Discount struct {
	gorm.Model
	FixedAmountID        *uint                       `json:"fixed_amount_id,omitempty"`
	FixedAmount          *FixedAmountDiscount        `json:"fixed_amount,omitempty" gorm:"foreignKey:FixedAmountID"`
	PercentageID         *uint                       `json:"percentage_id,omitempty"`
	Percentage           *PercentageDiscount         `json:"percentage,omitempty" gorm:"foreignKey:PercentageID"`
	PercentageCategoryID *uint                       `json:"percentage_category_id,omitempty"`
	PercentageCategory   *PercentageCategoryDiscount `json:"percentage_category,omitempty" gorm:"foreignKey:PercentageCategoryID"`
	SeasonalID           *uint                       `json:"seasonal_id,omitempty"`
	Seasonal             *SeasonalDiscount           `json:"seasonal,omitempty" gorm:"foreignKey:SeasonalID"`
	PointDiscountID      *uint                       `json:"point_discount_id,omitempty"`
	PointDiscount        *PointDiscount              `json:"point_discount,omitempty" gorm:"foreignKey:PointDiscountID"`
	DiscountCategoryID   uint                        `json:"discount_category_id"`
	DiscountCategory     DiscountCategory            `json:"discount_category" gorm:"foreignKey:DiscountCategoryID;unique"`
}
