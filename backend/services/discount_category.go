package services

import (
	"context"
	"playtorium/models"

	"gorm.io/gorm"
)

type DiscountCategoryService interface {
	GetAvailableDiscountCategories(ctx context.Context) ([]models.DiscountCategory, error)
}

type discountCategoryService struct {
	db *gorm.DB
}

func NewDiscountCategoryService(db *gorm.DB) DiscountCategoryService {
	return &discountCategoryService{db: db}
}

func (s *discountCategoryService) GetAvailableDiscountCategories(ctx context.Context) ([]models.DiscountCategory, error) {
	var discountCategories []models.DiscountCategory

	if err := s.db.Where("id NOT IN (SELECT discount_category_id FROM discounts)").
		Find(&discountCategories).Error; err != nil {
		return nil, err
	}

	return discountCategories, nil
}
