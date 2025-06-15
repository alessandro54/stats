package app

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"

	"github.com/alessandro54/stats/internal/adapters/blizzard"
)

func Start() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	app := fiber.New()

	// Example health check route
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	provider := blizzard.NewBlizzardTokenProvider()

	app.Get("/token", func(c fiber.Ctx) error {
		token, err := provider.GetToken()
		if err != nil {
			return c.Status(500).SendString("Failed to get token: " + err.Error())
		}
		return c.JSON(fiber.Map{"access_token": token})
	})

	// Get PORT from env or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server listening on http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
