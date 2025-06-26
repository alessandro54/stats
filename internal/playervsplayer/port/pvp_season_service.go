package port

import (
	"context"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
)

type PvpSeasonService interface {
	FetchOrInsert(ctx context.Context, blizzardID uint) (*model.PvpSeason, error)
}
