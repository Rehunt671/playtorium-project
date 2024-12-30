package handlers

import (
	"net/http"
	"strings"

	"playtorium/services"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetMe(c *gin.Context)
}

type userHandler struct {
	authService services.AuthService
	userService services.UserService
}

func NewUserHandler(authService services.AuthService, userService services.UserService) UserHandler {
	return &userHandler{
		authService: authService,
		userService: userService,
	}
}

func (h *userHandler) GetMe(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := h.authService.ExtractUserIDFromToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user details"})
		return
	}

	user.Password = ""
	c.JSON(http.StatusOK, user)
}
