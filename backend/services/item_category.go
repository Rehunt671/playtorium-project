package services

import (
	"playtorium/models"

	"gorm.io/gorm"
)

type ItemCategoryService interface {
	GetAllItemCategories() ([]models.ItemCategory, error)
}

type itemCategoryServiceImpl struct {
	db *gorm.DB
}

func NewItemCategoryService(db *gorm.DB) ItemCategoryService {
	return &itemCategoryServiceImpl{db: db}
}

func (s *itemCategoryServiceImpl) GetAllItemCategories() ([]models.ItemCategory, error) {
	var categories []models.ItemCategory
	if err := s.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
