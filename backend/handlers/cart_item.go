package handlers

import (
	"net/http"
	"playtorium/models"
	"playtorium/services"

	"github.com/gin-gonic/gin"
)

type CartItemHandler interface {
	AddCartItem(c *gin.Context)
	UpdateCartItem(c *gin.Context)
	RemoveCartItem(c *gin.Context)
}

type cartItemHandlerImpl struct {
	service services.CartItemService
}

func NewCartItemHandler(service services.CartItemService) CartItemHandler {
	return &cartItemHandlerImpl{service: service}
}

func (h *cartItemHandlerImpl) AddCartItem(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.AddCartItem(&cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func (h *cartItemHandlerImpl) UpdateCartItem(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateCartItem(&cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item updated in cart"})
}

func (h *cartItemHandlerImpl) RemoveCartItem(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.RemoveCartItem(&cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
