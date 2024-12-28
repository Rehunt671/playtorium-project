package models

import "gorm.io/gorm"

type DiscountCategory struct {
    gorm.Model
    Name string `json:"name" gorm:"unique"`
}