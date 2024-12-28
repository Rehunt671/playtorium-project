package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID   uint `json:"cart_id"`
	ItemID   uint `json:"item_id"`
	Item     Item `json:"item" gorm:"foreignKey:ItemID"`
	Quantity int  `json:"quantity"`
}
