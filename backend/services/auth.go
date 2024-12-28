package services

import (
	"errors"
	"playtorium/dtos"
	"playtorium/models"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(credentials *dtos.RegisterCredentials) error
	Login(credentials *dtos.LoginCredentials) (string, error)
}

type AuthServiceImpl struct {
	db          *gorm.DB
	cartService CartService
}

func NewAuthService(db *gorm.DB, cartService CartService) AuthService {
	return &AuthServiceImpl{db: db, cartService: cartService}
}

func (s *AuthServiceImpl) Register(credentials *dtos.RegisterCredentials) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	credentials.Password = string(hashedPassword)

	user := &models.User{
		Name:     credentials.Name,
		Username: credentials.Username,
		Password: credentials.Password,
		Points:   0,
	}

	cart := &models.Cart{User: *user}
	if err := s.cartService.CreateCart(cart); err != nil {
		return err
	}

	return nil
}

func (s *AuthServiceImpl) Login(credentials *dtos.LoginCredentials) (string, error) {
	var user models.User
	if err := s.db.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
		return "", errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
