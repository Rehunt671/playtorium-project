package models

import "gorm.io/gorm"

type Campaign struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}
