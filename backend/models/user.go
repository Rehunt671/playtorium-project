package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty" gorm:"unique"`
	Password string `json:"password,omitempty"`
	Points   int    `json:"points,omitempty" gorm:"default:0"`
	Cart     Cart   `json:"cart,omitempty" gorm:"foreignKey:UserID"`
}
