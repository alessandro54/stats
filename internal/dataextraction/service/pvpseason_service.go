package service

import (
	"context"
	"encoding/json"
	"fmt"
	pvpseason "github.com/alessandro54/stats/internal/dataextraction/adapter/blizzard/gamedata"
	port2 "github.com/alessandro54/stats/internal/dataextraction/port"
)

type pvpSeasonService struct {
	snapshotService port2.SnapshotService
}

func NewPvpService(snapshotService port2.SnapshotService) port2.PvpSeasonService {
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

func (p pvpSeasonService) GetCurrentSeasonID(ctx context.Context) (int, error) {
	data, err := pvpseason.FetchPvpSeasonIndex(ctx, map[string]string{})
	if err != nil {
		return 0, err
	}

	var result struct {
		CurrentSeason struct {
			ID int `json:"id"`
		} `json:"current_season"`
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("unmarshal error: %w", err)
	}

	return result.CurrentSeason.ID, nil
}
