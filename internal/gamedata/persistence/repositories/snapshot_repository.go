package repositories

import (
	"context"
	"errors"
	"github.com/alessandro54/stats/internal/gamedata/model"
	"github.com/alessandro54/stats/internal/gamedata/port"
	"gorm.io/gorm"
)

type snapshotRepositoryImpl struct {
	db *gorm.DB
}

func NewSnapshotRepository(db *gorm.DB) port.SnapshotRepository {
	return &snapshotRepositoryImpl{db: db}
}

func (r *snapshotRepositoryImpl) SaveSnapshot(ctx context.Context, snapshot *model.LeaderboardSnapshot) error {
	return r.db.WithContext(ctx).Create(snapshot).Error
}

func (r *snapshotRepositoryImpl) GetAllSnapshots(ctx context.Context) ([]*model.LeaderboardSnapshot, error) {
	var snapshots []*model.LeaderboardSnapshot
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&snapshots).Error
	if err != nil {
		return nil, err
	}
	return snapshots, nil
}

func (r *snapshotRepositoryImpl) GetLatestSnapshot(ctx context.Context, mode string, bracket string) (*model.LeaderboardSnapshot, error) {
	var snapshot model.LeaderboardSnapshot
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
