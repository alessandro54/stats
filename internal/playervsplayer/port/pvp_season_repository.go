package port

import (
	"context"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
)

type PvpSeasonRepository interface {
	FindByBlizzardID(ctx context.Context, blizzardID uint) (*model.PvpSeason, error)
	Insert(ctx context.Context, season *model.PvpSeason) error
}
