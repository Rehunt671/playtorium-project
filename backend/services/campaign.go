package services

import (
	"context"
	"playtorium/models"

	"gorm.io/gorm"
)

type CampaignService interface {
	GetCampaignsByDiscountCategoryId(ctx context.Context, discountCategoryId int) ([]models.Campaign, error)
}

type campaignServiceImpl struct {
	db *gorm.DB
}

func NewCampaignService(db *gorm.DB) CampaignService {
	return &campaignServiceImpl{db: db}
}

func (s *campaignServiceImpl) GetCampaignsByDiscountCategoryId(ctx context.Context, discountCategoryId int) ([]models.Campaign, error) {
	var campaignOnDiscountCategories []models.CampaignOnDiscountCategory
	var campaigns []models.Campaign

	if err := s.db.Where("discount_category_id = ?", discountCategoryId).
		Preload("Campaign").
		Find(&campaignOnDiscountCategories).Error; err != nil {
		return nil, err
	}

	for _, codc := range campaignOnDiscountCategories {
		campaigns = append(campaigns, codc.Campaign)
	}

	return campaigns, nil
}
