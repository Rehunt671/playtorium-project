package models

type PercentageCategoryDiscount struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	ItemCategoryID uint         `json:"item_category_id"`
	ItemCategory   ItemCategory `json:"item_category" gorm:"foreignKey:ItemCategoryID"`
	Percentage     float64      `json:"percentage"`
}
