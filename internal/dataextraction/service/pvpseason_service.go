package service

import (
	"context"
	"fmt"
	pvpseason "github.com/alessandro54/stats/internal/dataextraction/adapter/blizzard/gamedata"
	"github.com/alessandro54/stats/internal/dataextraction/port"
)

type pvpSeasonService struct {
	snapshotService port.SnapshotService
}

func NewPvpService(snapshotService port.SnapshotService) port.PvpSeasonService {
	return &pvpSeasonService{
		snapshotService: snapshotService,
	}
}

func (p pvpSeasonService) GetLatestPvpLeaderboard(ctx context.Context, bracket string) ([]byte, error) {
	data, err := p.snapshotService.GetLatestSeasonByBracket(ctx, "pvp", bracket)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch leaderboard: %w", err)
	}
	if data == nil {
		return nil, fmt.Errorf("no leaderboard data found for bracket: %s", bracket)
	}

	return data.Data, nil
}

func (p pvpSeasonService) GetCurrentSeasonID(ctx context.Context, region string) (int, error) {
	data, err := pvpseason.FetchPvpSeasonIndex(ctx, region)
	if err != nil {
		return 0, err
	}

	return data.CurrentSeason.ID, nil
}
