package handler

import (
	"github.com/alessandro54/stats/internal/gameinfo/domain/port"
	"github.com/gofiber/fiber/v3"
)

type LeaderboardSnapshotHandler struct {
	service port.LeaderboardSnapshotService
}

func NewLeaderboardSnapshotHandler(service port.LeaderboardSnapshotService) *LeaderboardSnapshotHandler {
	return &LeaderboardSnapshotHandler{
		service: service,
	}
}

func (h *LeaderboardSnapshotHandler) GetAllSnapshots(c fiber.Ctx) error {
	snapshots, err := h.service.GetAll(c.Context())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve snapshots",
		})
	}

	return c.Status(fiber.StatusOK).JSON(snapshots)
}
