package port

import "context"

type PvpSeasonService interface {
	GetCurrentSeasonID(ctx context.Context) (int, error)
	GetLatestPvpLeaderboard(ctx context.Context, bracket string) ([]byte, error)
}
