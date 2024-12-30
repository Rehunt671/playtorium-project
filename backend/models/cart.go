package models

type Cart struct {
	ID     uint       `json:"id" gorm:"primaryKey"`
	UserID uint       `json:"user_id"`
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID"`
}
