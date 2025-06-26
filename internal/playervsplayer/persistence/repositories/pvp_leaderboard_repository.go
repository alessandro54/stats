package repositories

import (
	"context"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
	"github.com/alessandro54/stats/internal/playervsplayer/port"
	"gorm.io/gorm"
)

type pvpLeaderboardRepositoryImpl struct {
	db *gorm.DB
}

func (p pvpLeaderboardRepositoryImpl) FindOrInsert(ctx context.Context, leaderboard *model.PvpLeaderboard) (*model.PvpLeaderboard, error) {
	panic("Implement me!")
}

func NewPvpLeaderboardRepository(db *gorm.DB) port.PvpLeaderboardRepository {
	return &pvpLeaderboardRepositoryImpl{
		db: db,
	}
}
