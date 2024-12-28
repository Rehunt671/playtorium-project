package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty" gorm:"unique"`
	Password string `json:"password,omitempty"`
	Points   int    `json:"points,omitempty" gorm:"default:0"`
	Carts    []Cart `json:"carts,omitempty" gorm:"foreignKey:UserID"`
}
