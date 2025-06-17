package main

import (
	"context"
	"github.com/alessandro54/stats/internal/gameinfo/adapter/blizzard"
	"github.com/alessandro54/stats/internal/shared"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()

	tokenProvider := blizzard.NewTokenProvider()

	app.Get("/debug/token", func(c fiber.Ctx) error {
		token, err := tokenProvider.GetToken(context.Background())
		if err != nil {
			return c.Status(500).SendString("Failed to get token: " + err.Error())
		}
		return c.JSON(fiber.Map{
			"token": token,
		})
	})

	// You would typically pass blizzardClient into your app layer here

	err = shared.StartFiberServer(shared.ServerConfig{
		Port:         8080,
		App:          app,
		GracefulWait: 5 * time.Second,
	})

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
