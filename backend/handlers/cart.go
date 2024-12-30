package handlers

import (
	"net/http"
	"playtorium/models"
	"playtorium/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartHandler interface {
	GetCartDetail(c *gin.Context)
	GetCartItemsByCartID(c *gin.Context)
	AddCartItem(c *gin.Context)
	UpdateCartItem(c *gin.Context)
	RemoveCartItem(c *gin.Context)
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

func (h *cartHandlerImpl) GetCartItemsByCartID(c *gin.Context) {
	cartIdStr := c.Param("cart_id")
	cartID, err := strconv.Atoi(cartIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
		return
	}

	cartItems, err := h.cartService.GetCartItemsByCartID(cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart items"})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}

func (h *cartHandlerImpl) AddCartItem(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.cartService.AddCartItem(&cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func (h *cartHandlerImpl) UpdateCartItem(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.cartService.UpdateCartItem(&cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item updated in cart"})
}

func (h *cartHandlerImpl) RemoveCartItem(c *gin.Context) {
	cartItemIdStr := c.Param("cart_item_id")
	cartItemId, err := strconv.Atoi(cartItemIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart ID"})
		return
	}

	err = h.cartService.RemoveCartItem(cartItemId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
