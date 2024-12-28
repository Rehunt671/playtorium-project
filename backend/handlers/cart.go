package handlers

import (
	"net/http"
	"playtorium/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartHandler interface {
	GetCartDetail(c *gin.Context)
}

type cartHandlerImpl struct {
	cartService services.CartService
}

func NewCartHandler(cartService services.CartService) CartHandler {
	return &cartHandlerImpl{cartService: cartService}
}

func (h *cartHandlerImpl) GetCartDetail(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	cartDetail, err := h.cartService.GetCartDetail(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, cartDetail)
}
