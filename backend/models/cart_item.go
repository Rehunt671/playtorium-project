package models

import "time"

type CartItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CartID    uint      `json:"cart_id"`
	ItemID    uint      `json:"item_id"`
	Item      Item      `json:"item" gorm:"foreignKey:ItemID"`
	Quantity  int       `json:"quantity"`
}
