package main

import (
	"github.com/alessandro54/stats/cmd/cron"
	"github.com/alessandro54/stats/infra/db"
	"github.com/alessandro54/stats/internal/gamedata/container"
	"github.com/alessandro54/stats/internal/system/handler"

	"github.com/alessandro54/stats/internal/shared"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Connect()
	db.RunMigrations(db.DB)

	app := fiber.New()

	api := app.Group("/api/v1")

	appContainer := container.InitAppContainer()

	handler.RegisterRoutes(api, appContainer)

	cron.StartCronJobs()

	err = shared.StartFiberServer(shared.ServerConfig{
		Port:         8080,
		App:          app,
		GracefulWait: 5 * time.Second,
	})

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
