package repositories

import (
	"context"
	client "github.com/alessandro54/stats/infra/db"
	"github.com/alessandro54/stats/internal/gameinfo/domain/entity"
	"gorm.io/gorm"
)

type LeaderboardSnapshotRepositoryImpl struct {
	db *gorm.DB
}

func NewLeaderboardSnapshotRepository() *LeaderboardSnapshotRepositoryImpl {
	return &LeaderboardSnapshotRepositoryImpl{db: client.DB}
}

func (r *LeaderboardSnapshotRepositoryImpl) SaveSnapshot(ctx context.Context, snapshot *entity.LeaderboardSnapshot) error {
	return r.db.WithContext(ctx).Create(snapshot).Error
}

func (r *LeaderboardSnapshotRepositoryImpl) GetAllSnapshots(ctx context.Context) ([]*entity.LeaderboardSnapshot, error) {
	var snapshots []*entity.LeaderboardSnapshot
	err := r.db.WithContext(ctx).Order("created_at DESC").Find(&snapshots).Error
	if err != nil {
		return nil, err
	}
	return snapshots, nil
}
