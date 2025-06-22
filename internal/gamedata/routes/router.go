package routes

import (
	"github.com/alessandro54/stats/internal/gamedata/container"
	"github.com/gofiber/fiber/v3"
)

func RegisterGameDataRoutes(router fiber.Router) {
	pvp := router.Group("/game-data/pvp")

	pvpSeasonHandler := container.InitPvpSeasonWire()

	pvp.Get("/current-season", pvpSeasonHandler.GetPvpSeasonID)
	pvp.Get("/leaderboard/:seasonId/:bracket", pvpSeasonHandler.GetPvpLeaderboard)
}
