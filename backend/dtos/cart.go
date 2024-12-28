package dtos

import "playtorium/models"

type CartDetail struct {
	Cart           models.Cart `json:"cart"`
	TotalPrice     float64     `json:"total_price"`
	UserPointsUsed int         `json:"user_points_used"`
}
