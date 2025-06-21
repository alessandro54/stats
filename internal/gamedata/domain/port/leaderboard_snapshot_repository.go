package port

import (
	"context"
	"github.com/alessandro54/stats/internal/gamedata/domain/entity"
)

type LeaderboardSnapshotRepository interface {
	GetAllSnapshots(ctx context.Context) ([]*entity.LeaderboardSnapshot, error)
	SaveSnapshot(ctx context.Context, snapshot *entity.LeaderboardSnapshot) error
}
