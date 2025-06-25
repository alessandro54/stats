package port

import (
	"context"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
)

type SnapshotService interface {
	GetAll(ctx context.Context) ([]*model.PvpLeaderboardSnapshot, error)
	Save(ctx context.Context, mode string, bracket string, data []byte) error

	GetLatestSeasonByBracket(ctx context.Context, mode string, bracket string) (*model.PvpLeaderboardSnapshot, error)
	FetchFromBlizzardAndSave(ctx context.Context, pvpSeasonId string, pvpBracket string, region string) error
}
