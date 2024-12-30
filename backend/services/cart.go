package services

import (
	"playtorium/dtos"
	"playtorium/models"

	"gorm.io/gorm"
)

type CartService interface {
	GetCartDetail(userID int) (*dtos.CartDetail, error)
	CreateCart(cart *models.Cart) error
	GetCartItemsByCartID(cartID int) ([]models.CartItem, error)
	AddCartItem(cartItem *models.CartItem) error
	UpdateCartItem(cartItem *models.CartItem) error
	RemoveCartItem(cartItemId int) error
}

type cartServiceImpl struct {
	db              *gorm.DB
	discountService DiscountService
}

func NewCartService(db *gorm.DB, discountService DiscountService) CartService {
	return &cartServiceImpl{db: db, discountService: discountService}
}

func (s *cartServiceImpl) GetCartDetail(userId int) (*dtos.CartDetail, error) {
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

func (s *cartServiceImpl) CreateCart(cart *models.Cart) error {
	if err := s.db.Create(&cart).Error; err != nil {
		return err
	}
	return nil
}

func (s *cartServiceImpl) GetCartItemsByCartID(cartId int) ([]models.CartItem, error) {
	var cartItems []models.CartItem

	err := s.db.Where("cart_id = ?", cartId).
		Joins("JOIN items ON cart_items.item_id = items.id").
		Order("items.name DESC").
		Preload("Item").
		Find(&cartItems).Error

	if err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (s *cartServiceImpl) AddCartItem(cartItem *models.CartItem) error {
	if err := s.db.Create(cartItem).Error; err != nil {
		return err
	}
	return nil
}

func (s *cartServiceImpl) UpdateCartItem(cartItem *models.CartItem) error {
	if err := s.db.Save(cartItem).Error; err != nil {
		return err
	}
	return nil
}

func (s *cartServiceImpl) RemoveCartItem(cartItemId int) error {
	if err := s.db.Delete(&models.CartItem{}, cartItemId).Error; err != nil {
		return err
	}
	return nil
}
