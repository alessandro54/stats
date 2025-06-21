package handler

import (
	"github.com/alessandro54/stats/internal/gamedata/handler"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(api fiber.Router) {
	api.Get("/health", healthHandler)

	handler.RegisterGameDataRoutes(api)
}

func healthHandler(c fiber.Ctx) error {
	return c.SendString("OK")
}
