package models

import "gorm.io/gorm"

type CampaignOnDiscountCategory struct {
	gorm.Model
	CampaignID         uint             `json:"campaign_id"`
	Campaign           Campaign         `json:"campaign" gorm:"foreignKey:CampaignID"`
	DiscountCategoryID uint             `json:"discount_category_id"`
	DiscountCategory   DiscountCategory `json:"discount_category" gorm:"foreignKey:DiscountCategoryID"`
}
