package port

import (
	"context"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
)

type SnapshotRepository interface {
	GetAllSnapshots(ctx context.Context) ([]*model.PvpLeaderboardSnapshot, error)
	GetLatestSnapshot(ctx context.Context, mode string, bracket string) (*model.PvpLeaderboardSnapshot, error)
	SaveSnapshot(ctx context.Context, snapshot *model.PvpLeaderboardSnapshot) error
}
