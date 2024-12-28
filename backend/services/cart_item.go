package services

import (
	"playtorium/models"

	"gorm.io/gorm"
)

type CartItemService interface {
	AddCartItem(cartItem *models.CartItem) error
	UpdateCartItem(cartItem *models.CartItem) error
	RemoveCartItem(cartItem *models.CartItem) error
}

type CartItemServiceImpl struct {
	db *gorm.DB
}

func NewCartItemService(db *gorm.DB) CartItemService {
	return &CartItemServiceImpl{db: db}
}

func (s *CartItemServiceImpl) AddCartItem(cartItem *models.CartItem) error {
	if err := s.db.Create(cartItem).Error; err != nil {
		return err
	}
	return nil
}

func (s *CartItemServiceImpl) UpdateCartItem(cartItem *models.CartItem) error {
	if err := s.db.Save(cartItem).Error; err != nil {
		return err
	}
	return nil
}

func (s *CartItemServiceImpl) RemoveCartItem(cartItem *models.CartItem) error {
	if err := s.db.Delete(cartItem).Error; err != nil {
		return err
	}
	return nil
}
