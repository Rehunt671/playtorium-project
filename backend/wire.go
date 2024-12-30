//go:build wireinject
// +build wireinject

package main

import (
	"playtorium/db"
	"playtorium/handlers"
	"playtorium/services"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeApp() (*gin.Engine, func(), error) {
	wire.Build(
		AppSet,
		HandlerSet,
		ServiceSet)

	return gin.New(), func() {}, nil
}

var AppSet = wire.NewSet(
	NewApp,
	db.NewPostgresDatabase,
)

var HandlerSet = wire.NewSet(
	handlers.NewAuthHandler,
	handlers.NewCartHandler,
	handlers.NewCampaignHandler,
	handlers.NewDiscountHandler,
	handlers.NewDiscountCategoryHandler,
	handlers.NewItemHandler,
	handlers.NewItemCategoryHandler,
	handlers.NewUserHandler,
)

var ServiceSet = wire.NewSet(
	services.NewAuthService,
	services.NewCartService,
	services.NewCampaignService,
	services.NewDiscountService,
	services.NewDiscountCategoryService,
	services.NewItemService,
	services.NewItemCategoryService,
	services.NewUserService,
)
