package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID uint       `json:"user_id"`
	User   User       `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID"`
}
