package handler

import "github.com/gofiber/fiber/v3"

func RegisterGameDataRoutes(router fiber.Router) {
	pvp := router.Group("/gamedata/pvp")

	pvp.Get("/current-season", GetPvpSeasonID)
	pvp.Get("/leaderboard/:seasonId/:bracket", GetPvpLeaderboard)
}
