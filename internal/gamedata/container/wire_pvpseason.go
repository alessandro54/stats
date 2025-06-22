//go:build wireinject
// +build wireinject

package container

import (
	"github.com/alessandro54/stats/internal/gamedata/handler"
	"github.com/alessandro54/stats/internal/gamedata/persistence/repositories"
	"github.com/alessandro54/stats/internal/gamedata/service"
	"github.com/google/wire"
)

func InitPvpSeasonWire() *handler.PvpSeasonHandler {
	wire.Build(
		repositories.NewLeaderboardSnapshotRepository,
		service.NewSnapshotService,
		service.NewPvpService,
		handler.NewPvpSeasonHandler)
	return &handler.PvpSeasonHandler{}
}
