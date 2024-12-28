package models

type Cart struct {
	ID     uint       `json:"id" gorm:"primaryKey"`
	UserID uint       `json:"user_id"`
	User   User       `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID"`
}
