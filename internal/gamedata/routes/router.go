package routes

import (
	"github.com/alessandro54/stats/internal/gamedata/container"
	"github.com/gofiber/fiber/v3"
)

func RegisterGameDataRoutes(router fiber.Router, c *container.AppContainer) {
	pvp := router.Group("/game-data/pvp")

	pvp.Get("/current-season", c.PvpSeasonHandler.GetPvpSeasonID)
	pvp.Get("/leaderboard/:seasonId/:bracket", c.PvpSeasonHandler.GetPvpLeaderboard)
}
