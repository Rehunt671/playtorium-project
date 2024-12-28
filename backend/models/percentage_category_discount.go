package models

import "gorm.io/gorm"

type PercentageCategoryDiscount struct {
	gorm.Model
	ItemCategoryID uint         `json:"item_category_id"`
	ItemCategory   ItemCategory `json:"item_category" gorm:"foreignKey:ItemCategoryID"`
	Percentage     float64      `json:"percentage"`
}
