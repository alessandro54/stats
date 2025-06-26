package repositories

import (
	"context"
	"errors"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
	"github.com/alessandro54/stats/internal/playervsplayer/port"
	"gorm.io/gorm"
)

type pvpSeasonRepositoryImpl struct {
	db *gorm.DB
}

func (p *pvpSeasonRepositoryImpl) FindByBlizzardID(ctx context.Context, blizzardID uint) (*model.PvpSeason, error) {
	var season model.PvpSeason
	err := p.db.WithContext(ctx).
		Where("blizzard_id = ?", blizzardID).
		First(&season).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &season, err
}

func (p *pvpSeasonRepositoryImpl) Insert(ctx context.Context, season *model.PvpSeason) error {
	return p.tryInsert(ctx, season)
}

func (p *pvpSeasonRepositoryImpl) tryInsert(ctx context.Context, season *model.PvpSeason) error {
	return p.db.WithContext(ctx).Create(season).Error
}

func NewPvpSeasonRepository(db *gorm.DB) port.PvpSeasonRepository {
	return &pvpSeasonRepositoryImpl{db: db}
}
