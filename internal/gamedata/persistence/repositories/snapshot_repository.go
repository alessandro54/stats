package repositories

import (
	"context"
	"errors"
	client "github.com/alessandro54/stats/infra/db"
	"github.com/alessandro54/stats/internal/gamedata/domain/entity"
	"github.com/alessandro54/stats/internal/gamedata/domain/port"
	"gorm.io/gorm"
)

type snapshotRepositoryImpl struct {
	db *gorm.DB
}

func NewLeaderboardSnapshotRepository() port.SnapshotRepository {
	return &snapshotRepositoryImpl{db: client.DB}
}

func (r *snapshotRepositoryImpl) SaveSnapshot(ctx context.Context, snapshot *entity.LeaderboardSnapshot) error {
	return r.db.WithContext(ctx).Create(snapshot).Error
}

func (r *snapshotRepositoryImpl) GetAllSnapshots(ctx context.Context) ([]*entity.LeaderboardSnapshot, error) {
	var snapshots []*entity.LeaderboardSnapshot
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&snapshots).Error
	if err != nil {
		return nil, err
	}
	return snapshots, nil
}

func (r *snapshotRepositoryImpl) GetLatestSnapshot(ctx context.Context, mode string, bracket string) (*entity.LeaderboardSnapshot, error) {
	var snapshot entity.LeaderboardSnapshot
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
