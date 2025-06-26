package handler

import (
	"encoding/json"
	"github.com/alessandro54/stats/internal/dataextraction/port"
	"github.com/gofiber/fiber/v3"
)

type PvpSeasonHandler struct {
	svc port.PvpSeasonService
}

func NewPvpSeasonHandler(s port.PvpSeasonService) *PvpSeasonHandler {
	return &PvpSeasonHandler{
		svc: s,
	}
}

func (h *PvpSeasonHandler) GetPvpLeaderboard(c fiber.Ctx) error {
	bracket := c.Params("bracket")
	snapshot, err := h.svc.GetLatestPvpLeaderboard(c.Context(), bracket)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch leaderboard",
		})
	}
	if snapshot == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No leaderboard data found for the specified bracket",
		})
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(snapshot, &parsed); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse leaderboard data",
		})
	}

	return c.JSON(fiber.Map{
		"leaderboard": parsed,
	})
}

func (h *PvpSeasonHandler) GetPvpSeasonID(c fiber.Ctx) error {
	id, err := h.svc.GetCurrentSeasonID(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch current season ID",
		})
	}

	return c.JSON(fiber.Map{
		"current_season_id": id,
	})
}
