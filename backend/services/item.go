package services

import (
	"playtorium/models"

	"gorm.io/gorm"
)

type ItemService interface {
	GetAllItems() ([]models.Item, error)
}

type itemService struct {
	db *gorm.DB
}

func NewItemService(db *gorm.DB) ItemService {
	return &itemService{db: db}
}

func (s *itemService) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	if err := s.db.Preload("ItemCategory").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
