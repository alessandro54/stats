package repositories

import (
	"context"
	"errors"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
	"github.com/alessandro54/stats/internal/playervsplayer/port"
	"gorm.io/gorm"
)

type pvpLeaderboardRepositoryImpl struct {
	db *gorm.DB
}

func (p *pvpLeaderboardRepositoryImpl) FindBySeasonAndBracket(ctx context.Context, seasonID uint, bracket string, region string) (*model.PvpLeaderboard, error) {
	var leaderboard model.PvpLeaderboard

	err := p.db.WithContext(ctx).
		Where("pvp_season_id = ? AND bracket = ? AND region = ?", seasonID, bracket, region).
		First(&leaderboard).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &leaderboard, err
}

func (p *pvpLeaderboardRepositoryImpl) Insert(ctx context.Context, leaderboard *model.PvpLeaderboard) error {
	return p.db.WithContext(ctx).Create(leaderboard).Error
}

func NewPvpLeaderboardRepository(db *gorm.DB) port.PvpLeaderboardRepository {
	return &pvpLeaderboardRepositoryImpl{
		db: db,
	}
}
