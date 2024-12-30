package services

import (
	"errors"
	"playtorium/models"

	"gorm.io/gorm"
)

type UserService interface {
	GetUserByID(userID uint) (*models.User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := s.db.Preload("Cart").First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
