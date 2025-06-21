package main

import (
	"github.com/alessandro54/stats/cmd/cron"
	"github.com/alessandro54/stats/infra/db"
	"github.com/alessandro54/stats/internal/gamedata/handler"
	"github.com/alessandro54/stats/internal/gamedata/persistence/repositories"
	"github.com/alessandro54/stats/internal/gamedata/services"
	"github.com/alessandro54/stats/internal/shared"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	app := fiber.New()

	api := app.Group("/api/v1")

	err := godotenv.Load()

	db.Connect()
	db.RunMigrations(db.DB)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repo := repositories.NewLeaderboardSnapshotRepository()
	svc := services.NewSnapshotService(repo)
	snapshotHandler := handler.NewLeaderboardSnapshotHandler(svc)
	cron.StartCronJobs()

	api.Get("/snapshots", snapshotHandler.GetAllSnapshots)

	err = shared.StartFiberServer(shared.ServerConfig{
		Port:         8080,
		App:          app,
		GracefulWait: 5 * time.Second,
	})

	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
