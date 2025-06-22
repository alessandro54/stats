package handler

import (
	"github.com/alessandro54/stats/internal/gamedata/port"
	"github.com/gofiber/fiber/v3"
)

type SnapshotHandler struct {
	service port.SnapshotService
}

func NewLeaderboardSnapshotHandler(service port.SnapshotService) *SnapshotHandler {
	return &SnapshotHandler{
		service: service,
	}
}

func (h *SnapshotHandler) GetAllSnapshots(c fiber.Ctx) error {
	snapshots, err := h.service.GetAll(c.Context())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve snapshots",
		})
	}

	return c.Status(fiber.StatusOK).JSON(snapshots)
}
