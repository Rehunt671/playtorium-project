package models

type Discount struct {
	ID                   uint                        `json:"id" gorm:"primaryKey"`
	FixedAmountID        *uint                       `json:"fixed_amount_id,omitempty"`
	FixedAmount          *FixedAmountDiscount        `json:"fixed_amount,omitempty" gorm:"foreignKey:FixedAmountID;constraint:OnDelete:CASCADE;"`
	PercentageID         *uint                       `json:"percentage_id,omitempty"`
	Percentage           *PercentageDiscount         `json:"percentage,omitempty" gorm:"foreignKey:PercentageID;constraint:OnDelete:CASCADE;"`
	PercentageCategoryID *uint                       `json:"percentage_category_id,omitempty"`
	PercentageCategory   *PercentageCategoryDiscount `json:"percentage_category,omitempty" gorm:"foreignKey:PercentageCategoryID;constraint:OnDelete:CASCADE;"`
	SeasonalID           *uint                       `json:"seasonal_id,omitempty"`
	Seasonal             *SeasonalDiscount           `json:"seasonal,omitempty" gorm:"foreignKey:SeasonalID;constraint:OnDelete:CASCADE;"`
	PointDiscountID      *uint                       `json:"point_discount_id,omitempty"`
	PointDiscount        *PointDiscount              `json:"point_discount,omitempty" gorm:"foreignKey:PointDiscountID;constraint:OnDelete:CASCADE;"`
	DiscountCategoryID   uint                        `json:"discount_category_id" gorm:"unique"`
	DiscountCategory     DiscountCategory            `json:"discount_category" gorm:"foreignKey:DiscountCategoryID;constraint:OnDelete:CASCADE;"`
}
