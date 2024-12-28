package models

import "gorm.io/gorm"

type ItemCategory struct {
	gorm.Model
	Name string `json:"name,omitempty"`
}
