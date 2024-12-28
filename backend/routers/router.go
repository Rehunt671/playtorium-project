package routers

import (
	"net/http"
	"playtorium/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, authHandler handlers.AuthHandler, cartHandler handlers.CartHandler, cartItemHandler handlers.CartItemHandler) {
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the api v1",
		})
	})
	router := r.Group("/api")
	SetupAuthRouter(router, authHandler)
	SetupCartRouter(router, cartHandler)
	SetupCartItemRouter(router, cartItemHandler)
}
