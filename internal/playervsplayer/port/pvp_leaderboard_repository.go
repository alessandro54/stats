package port

import (
	"context"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
)

type PvpLeaderboardRepository interface {
	Insert(ctx context.Context, leaderboard *model.PvpLeaderboard) error
	FindBySeasonAndBracket(ctx context.Context, seasonID uint, bracket string, region string) (*model.PvpLeaderboard, error)
}
