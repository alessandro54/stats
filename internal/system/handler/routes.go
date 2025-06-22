package handler

import (
	"github.com/alessandro54/stats/internal/gamedata/container"
	"github.com/alessandro54/stats/internal/gamedata/routes"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(api fiber.Router, appContainer *container.AppContainer) {
	api.Get("/health", healthHandler)

	routes.RegisterGameDataRoutes(api, appContainer)
}

func healthHandler(c fiber.Ctx) error {
	return c.SendString("OK")
}
