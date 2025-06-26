package repositories

import (
	"context"
	"errors"
	"github.com/alessandro54/stats/internal/dataextraction/port"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
	"gorm.io/gorm"
)

type pvpSnapshotRepositoryImpl struct {
	db *gorm.DB
}

func NewSnapshotRepository(db *gorm.DB) port.SnapshotRepository {
	return &pvpSnapshotRepositoryImpl{db: db}
}

func (r *pvpSnapshotRepositoryImpl) SaveSnapshot(ctx context.Context, snapshot *model.PvpLeaderboardSnapshot) error {
	return r.db.WithContext(ctx).Create(snapshot).Error
}

func (r *pvpSnapshotRepositoryImpl) GetAllSnapshots(ctx context.Context) ([]*model.PvpLeaderboardSnapshot, error) {
	var snapshots []*model.PvpLeaderboardSnapshot
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&snapshots).Error
	if err != nil {
		return nil, err
	}
	return snapshots, nil
}

func (r *pvpSnapshotRepositoryImpl) GetLatestSnapshot(ctx context.Context, mode string, bracket string) (*model.PvpLeaderboardSnapshot, error) {
	var snapshot model.PvpLeaderboardSnapshot
	err := r.db.WithContext(ctx).
		Where("mode = ?", mode).
		Where("bracket = ?", bracket).
		Order("created_at DESC").
		First(&snapshot).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No snapshot found
		}
		return nil, err
	}
	return &snapshot, nil
}
