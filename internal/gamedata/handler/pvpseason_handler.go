package handler

import (
	"github.com/alessandro54/stats/internal/gamedata/service"
	"github.com/gofiber/fiber/v3"
)

func GetPvpLeaderboard(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "This endpoint is not implemented yet",
	})
}

func GetPvpSeasonID(c fiber.Ctx) error {
	id, err := service.GetCurrentSeasonID(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch current season ID",
		})
	}

	return c.JSON(fiber.Map{
		"current_season_id": id,
	})
}
