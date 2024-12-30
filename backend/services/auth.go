package services

import (
	"errors"
	"os"
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
	ExtractUserIDFromToken(token string) (uint, error)
}

type authServiceImpl struct {
	db          *gorm.DB
	cartService CartService
}

func NewAuthService(db *gorm.DB, cartService CartService) AuthService {
	return &authServiceImpl{db: db, cartService: cartService}
}
func (s *authServiceImpl) Register(credentials *dtos.RegisterCredentials) error {
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
		Cart:     models.Cart{},
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		user.Cart.UserID = user.ID
		if err := tx.Create(&user.Cart).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

func (s *authServiceImpl) Login(credentials *dtos.LoginCredentials) (string, error) {
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

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *authServiceImpl) ExtractUserIDFromToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		return userID, nil
	}

	return 0, errors.New("invalid token")
}
