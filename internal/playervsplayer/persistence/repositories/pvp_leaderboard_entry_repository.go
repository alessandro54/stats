package repositories

import (
	"context"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
	"github.com/alessandro54/stats/internal/playervsplayer/port"
	"gorm.io/gorm"
)

type pvpLeaderboardEntryRepositoryImpl struct {
	db *gorm.DB
}

func (r *pvpLeaderboardEntryRepositoryImpl) FindByLeaderboardID(ctx context.Context, leaderboardID uint) ([]model.PvpLeaderboardEntry, error) {
	var entries []model.PvpLeaderboardEntry
	err := r.db.WithContext(ctx).
		Where("pvp_leaderboard_id = ?", leaderboardID).
		Find(&entries).Error
	return entries, err
}

func (r *pvpLeaderboardEntryRepositoryImpl) FindOneByCharacterAndLeaderboard(ctx context.Context, characterID uint, leaderboardID uint) (*model.PvpLeaderboardEntry, error) {
	var entry model.PvpLeaderboardEntry
	err := r.db.WithContext(ctx).
		Where("character_id = ? AND pvp_leaderboard_id = ?", characterID, leaderboardID).
		First(&entry).Error
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *pvpLeaderboardEntryRepositoryImpl) Save(ctx context.Context, entry *model.PvpLeaderboardEntry) error {
	return r.db.WithContext(ctx).Save(entry).Error
}

func (r *pvpLeaderboardEntryRepositoryImpl) BulkInsert(ctx context.Context, entries []model.PvpLeaderboardEntry) error {
	return r.db.WithContext(ctx).Create(&entries).Error
}

func (r *pvpLeaderboardEntryRepositoryImpl) DeleteByLeaderboardID(ctx context.Context, leaderboardID uint) error {
	return r.db.WithContext(ctx).
		Where("pvp_leaderboard_id = ?", leaderboardID).
		Delete(&model.PvpLeaderboardEntry{}).Error
}

func NewPvpLeaderboardEntryRepository(db *gorm.DB) port.PvpLeaderboardEntryRepository {
	return &pvpLeaderboardEntryRepositoryImpl{db: db}
}
