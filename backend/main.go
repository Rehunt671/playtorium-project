package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"playtorium/configs"
	"playtorium/handlers"
	"playtorium/routers"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewApp(
	authHandler handlers.AuthHandler,
	cartHandler handlers.CartHandler,
	campaignHandler handlers.CampaignHandler,
	discountHandler handlers.DiscountHandler,
	discountCategoryHandler handlers.DiscountCategoryHandler,
	itemHandler handlers.ItemHandler,
	itemCategoryHandler handlers.ItemCategoryHandler,
	userHandler handlers.UserHandler,
) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // Allow content-type and authorization headers
		AllowCredentials: true,
	}))
	routers.SetupRoutes(r, authHandler, cartHandler, campaignHandler, discountHandler, discountCategoryHandler, itemHandler, itemCategoryHandler, userHandler)
	return r, nil
}

func main() {
	configs.InitialEnv(".env")
	app, cleanup, err := InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize the app: %v", err)
	}

	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exitChan
		log.Print("Shutting down the server...")
		cleanup()
		os.Exit(0)
	}()

	log.Printf("Starting the server on port %s...", configs.GetPort())
	if err := app.Run(fmt.Sprintf(":%s", configs.GetPort())); err != nil {
		log.Printf("Server error: %v", err)
		cleanup()
	}
}
