package service

import (
	"context"
	"encoding/json"
	"fmt"
	pvpseason "github.com/alessandro54/stats/internal/gamedata/adapter/blizzard/gamedata"
)

func GetCurrentSeasonID(ctx context.Context) (int, error) {
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
