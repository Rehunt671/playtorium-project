package models

type Item struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	Name           string       `json:"name,omitempty"`
	Price          float64      `json:"price,omitempty"`
	ItemCategoryID uint         `json:"item_category_id,omitempty"`
	ItemCategory   ItemCategory `json:"item_category,omitempty" gorm:"foreignKey:ItemCategoryID"`
	Stock          int          `json:"stock,omitempty"`
}
