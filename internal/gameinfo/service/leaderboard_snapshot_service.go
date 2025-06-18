package service

import (
	"context"
	"github.com/alessandro54/stats/internal/gameinfo/domain/entity"
	"github.com/alessandro54/stats/internal/gameinfo/domain/port"
	"time"
)

type snapshotService struct {
	repo port.LeaderboardSnapshotRepository
}

func NewSnapshotService(repo port.LeaderboardSnapshotRepository) port.LeaderboardSnapshotService {
	return &snapshotService{
		repo: repo,
	}
}

func (s *snapshotService) GetAll(ctx context.Context) ([]*entity.LeaderboardSnapshot, error) {
	return s.repo.GetAllSnapshots(ctx)
}

func (s *snapshotService) Save(ctx context.Context, mode string, bracket string, data []byte) error {
	snapshot := &entity.LeaderboardSnapshot{
		Mode:      mode,
		Bracket:   bracket,
		Data:      data,
		CreatedAt: time.Now(),
	}
	return s.repo.SaveSnapshot(ctx, snapshot)
}
