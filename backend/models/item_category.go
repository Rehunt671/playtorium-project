package models

type ItemCategory struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name,omitempty"`
}
