package services

import (
	"playtorium/dtos"
	"playtorium/models"

	"gorm.io/gorm"
)

type CartService interface {
	GetCartDetail(userID int) (*dtos.CartDetail, error)
	CreateCart(cart *models.Cart) error
}

type CartServiceImpl struct {
	db              *gorm.DB
	discountService DiscountService
}

func NewCartService(db *gorm.DB, discountService DiscountService) CartService {
	return &CartServiceImpl{db: db, discountService: discountService}
}

func (s *CartServiceImpl) GetCartDetail(userId int) (*dtos.CartDetail, error) {
	var cart models.Cart
	if err := s.db.Preload("Items.Item.ItemCategory").Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return nil, err
	}
	totalPrice, userPointsUsed, err := s.discountService.ApplyDiscounts(cart.Items)
	if err != nil {
		return nil, err
	}

	cartDetail := &dtos.CartDetail{
		Cart:           cart,
		TotalPrice:     totalPrice,
		UserPointsUsed: userPointsUsed,
	}

	return cartDetail, nil
}

func (s *CartServiceImpl) CreateCart(cart *models.Cart) error {
	if err := s.db.Create(&cart).Error; err != nil {
		return err
	}
	return nil
}
