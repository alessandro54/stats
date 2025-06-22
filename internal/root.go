package internal

import (
	"github.com/alessandro54/stats/cmd/cron"
	"github.com/alessandro54/stats/internal/common"
	"github.com/alessandro54/stats/internal/gamedata/container"
	"github.com/alessandro54/stats/internal/gamedata/routes"
	"github.com/alessandro54/stats/internal/infra/db"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func RegisterRoutes(api fiber.Router, appContainer *container.AppContainer) {
	api.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("OK")
	})

	routes.RegisterGameDataRoutes(api, appContainer)
}

func SetupApp() *fiber.App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err = db.Connect()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	app := fiber.New()
	api := app.Group("/api/v1")
	appContainer := container.InitAppContainer()

	RegisterRoutes(api, appContainer)

	cron.StartCronJobs()

	err = common.StartFiberServer(common.ServerConfig{
		Port:         8080,
		App:          app,
		GracefulWait: 5 * time.Second,
	})

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	return app
}
