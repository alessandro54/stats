//go:build wireinject
// +build wireinject

package container

import (
	"github.com/alessandro54/stats/internal/gamedata/handler"
	"github.com/alessandro54/stats/internal/gamedata/persistence/repositories"
	"github.com/alessandro54/stats/internal/gamedata/service"
	"github.com/alessandro54/stats/internal/infra/db"
	"github.com/google/wire"
)

// AppContainer holds all your initialized handlers
type AppContainer struct {
	PvpSeasonHandler *handler.PvpSeasonHandler
	// Add more handlers here as needed
}

func InitAppContainer() *AppContainer {
	wire.Build(
		// DB provider
		db.ProvideDB,

		// Repositories
		repositories.NewSnapshotRepository,

		// Services
		service.NewSnapshotService,
		service.NewPvpService,

		// Handlers
		handler.NewPvpSeasonHandler,

		// The actual container
		wire.Struct(new(AppContainer), "*"),
	)

	// Dummy return for wire
	return &AppContainer{}
}
