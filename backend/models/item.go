package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name           string       `json:"name,omitempty"`
	Price          float64      `json:"price,omitempty"`
	ItemCategoryID uint         `json:"item_category_id,omitempty"`
	ItemCategory   ItemCategory `json:"item_category,omitempty" gorm:"foreignKey:ItemCategoryID"`
	Stock          int          `json:"stock,omitempty"`
}
