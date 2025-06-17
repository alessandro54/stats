package main

import (
	"encoding/json"
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

	app.Get("/debug/equipment/:realm/:char", func(c fiber.Ctx) error {
		blizz := blizzard.GetClient()

		data, err := blizz.FetchCharacterEquipment(c.Context(), "illidan", "nystinn")

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		var result any
		if err := json.Unmarshal(data, &result); err != nil {
			return c.Status(500).SendString("failed to parse Blizzard JSON: " + err.Error())
		}
		return c.JSON(fiber.Map{"equipment": result})
	})

	app.Get("/debug/leaderboard/:pvpSeasonId/:pvpBracket", func(c fiber.Ctx) error {
		blizz := blizzard.GetClient()

		data, err := blizz.FetchPvpLeaderboard(c.Context(), c.Params("pvpSeasonId"), c.Params("pvpBracket"))

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		var result any
		if err := json.Unmarshal(data, &result); err != nil {
			return c.Status(500).SendString("failed to parse Blizzard JSON: " + err.Error())
		}
		return c.JSON(fiber.Map{"equipment": result})
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
