package repositories

import (
	"context"
	"errors"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
	"github.com/alessandro54/stats/internal/playervsplayer/port"
	"gorm.io/gorm"
	"strings"
)

type pvpSeasonRepositoryImpl struct {
	db *gorm.DB
}

func (p *pvpSeasonRepositoryImpl) FindByBlizzardID(ctx context.Context, blizzardID string) (*model.PvpSeason, error) {
	var season model.PvpSeason
	err := p.db.WithContext(ctx).
		Where("blizzard_id = ?", blizzardID).
		First(&season).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &season, err
}

func (p *pvpSeasonRepositoryImpl) FindOrInsert(ctx context.Context, season *model.PvpSeason) (*model.PvpSeason, error) {
	if existing, err := p.FindByBlizzardID(ctx, season.BlizzardID); err != nil {
		return nil, err
	} else if existing != nil {
		return existing, nil
	}

	if err := p.tryInsert(ctx, season); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) ||
			strings.Contains(err.Error(), "duplicate") {
			// someone else won the race – re-read
			return p.FindByBlizzardID(ctx, season.BlizzardID)
		}
		return nil, err
	}

	return season, nil // freshly inserted, ID now populated
}

func (p *pvpSeasonRepositoryImpl) Insert(ctx context.Context, season *model.PvpSeason) error {
	return p.tryInsert(ctx, season) // delegate
}

func (p *pvpSeasonRepositoryImpl) tryInsert(ctx context.Context, season *model.PvpSeason) error {
	return p.db.WithContext(ctx).Create(season).Error
}

func NewPvpSeasonRepository(db *gorm.DB) port.PvpSeasonRepository {
	return &pvpSeasonRepositoryImpl{db: db}
}
