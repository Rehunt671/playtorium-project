package models

type CampaignOnDiscountCategory struct {
	ID                 uint             `json:"id" gorm:"primaryKey"`
	CampaignID         uint             `json:"campaign_id"`
	Campaign           Campaign         `json:"campaign" gorm:"foreignKey:CampaignID"`
	DiscountCategoryID uint             `json:"discount_category_id"`
	DiscountCategory   DiscountCategory `json:"discount_category" gorm:"foreignKey:DiscountCategoryID"`
}
