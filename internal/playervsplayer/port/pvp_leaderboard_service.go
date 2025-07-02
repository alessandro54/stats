package port

import (
	"context"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
)

type PvpLeaderboardService interface {
	GetOrFetch(ctx context.Context, seasonBlizzardID uint, bracket string, region string) (*model.PvpLeaderboard, error)
}
