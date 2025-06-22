package port

import (
	"context"
	"github.com/alessandro54/stats/internal/gamedata/model"
)

type SnapshotRepository interface {
	GetAllSnapshots(ctx context.Context) ([]*model.LeaderboardSnapshot, error)
	GetLatestSnapshot(ctx context.Context, mode string, bracket string) (*model.LeaderboardSnapshot, error)
	SaveSnapshot(ctx context.Context, snapshot *model.LeaderboardSnapshot) error
}
