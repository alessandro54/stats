package port

import (
	"context"
	"github.com/alessandro54/stats/internal/gamedata/domain/entity"
)

type SnapshotRepository interface {
	GetAllSnapshots(ctx context.Context) ([]*entity.LeaderboardSnapshot, error)
	GetLatestSnapshot(ctx context.Context, mode string, bracket string) (*entity.LeaderboardSnapshot, error)
	SaveSnapshot(ctx context.Context, snapshot *entity.LeaderboardSnapshot) error
}
