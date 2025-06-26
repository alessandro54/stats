package port

import (
	"context"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
)

type PvpLeaderboardRepository interface {
	FindOrInsert(ctx context.Context, leaderboard *model.PvpLeaderboard) (*model.PvpLeaderboard, error)
}
