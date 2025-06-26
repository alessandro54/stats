package port

import (
	"context"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
)

type PvpLeaderboardEntryRepository interface {
	FindByLeaderboardID(ctx context.Context, leaderboardID uint) ([]model.PvpLeaderboardEntry, error)
	FindOneByCharacterAndLeaderboard(ctx context.Context, characterID uint, leaderboardID uint) (*model.PvpLeaderboardEntry, error)
	Save(ctx context.Context, entry *model.PvpLeaderboardEntry) error
	BulkInsert(ctx context.Context, entries []model.PvpLeaderboardEntry) error
	DeleteByLeaderboardID(ctx context.Context, leaderboardID uint) error
}
