package port

import (
	"context"
	"github.com/alessandro54/stats/internal/gamedata/domain/entity"
)

type SnapshotService interface {
	GetAll(ctx context.Context) ([]*entity.LeaderboardSnapshot, error)
	Save(ctx context.Context, mode string, bracket string, data []byte) error

	GetLatestSeasonByBracket(ctx context.Context, mode string, bracket string) (*entity.LeaderboardSnapshot, error)
	FetchFromBlizzardAndSave(ctx context.Context, pvpSeasonId string, pvpBracket string, region string) error
}
