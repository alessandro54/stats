package port

import (
	"context"
	"github.com/alessandro54/stats/internal/gameinfo/domain/entity"
)

type LeaderboardSnapshotService interface {
	GetAll(ctx context.Context) ([]*entity.LeaderboardSnapshot, error)
	Save(ctx context.Context, mode string, bracket string, data []byte) error

	FetchFromBlizzardAndSave(ctx context.Context, pvpSeasonId string, pvpBracket string) error
}
