package pvpseason

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alessandro54/stats/internal/dataextraction/adapter/blizzard"
	"github.com/alessandro54/stats/internal/dataextraction/adapter/blizzard/gamedata/response"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
	"time"
)

func FetchPvpSeasonIndex(ctx context.Context, region string) (*response.PvpSeasonIndexResponse, error) {
	client, _ := blizzard.GetClient(ctx, region, "en_US")

	raw, err := client.Get(
		ctx,
		"/data/wow/pvp-season/index",
		map[string]string{
			"namespace": "dynamic-" + client.Region,
			"locale":    client.Locale,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("blizzard fetch failed: %w", err)
	}

	var dto response.PvpSeasonIndexResponse
	if err := json.Unmarshal(raw, &dto); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return &dto, nil
}

func FetchPvpSeason(ctx context.Context, blizzardID uint, region string) (*model.PvpSeason, error) {
	client, _ := blizzard.GetClient(ctx, region, "en_US")

	raw, err := client.Get(
		ctx,
		fmt.Sprintf("/data/wow/pvp-season/%d", blizzardID),
		map[string]string{
			"namespace": "dynamic-" + client.Region,
			"locale":    client.Locale,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("blizzard fetch failed: %w", err)
	}

	var dto response.PvpSeasonResponse
	if err := json.Unmarshal(raw, &dto); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	start := time.UnixMilli(dto.SeasonStartUnixMs).UTC()
	var end *time.Time

	if dto.SeasonEndUnixMs > 0 {
		t := time.UnixMilli(dto.SeasonEndUnixMs).UTC()
		end = &t
	}

	return &model.PvpSeason{
		BlizzardID: dto.ID,
		Name:       dto.SeasonName,
		StartTime:  start,
		EndTime:    end,
	}, nil
}

func FetchLeaderboardIndex(ctx context.Context, seasonId string, opts map[string]string) ([]byte, error) {
	client, _ := blizzard.GetClient(ctx, opts["region"], opts["locale"])
	return client.Get(
		ctx,
		"/data/wow/pvp-season/"+seasonId+"/pvp-leaderboard/index",
		map[string]string{
			"namespace": "dynamic-" + client.Region,
			"locale":    client.Locale,
		})

}

func FetchLeaderboard(ctx context.Context, seasonId uint, bracket string, opts map[string]string) (*response.PvpLeaderboardResponse, error) {
	if opts == nil {
		opts = map[string]string{
			"region": "us",
			"locale": "en_US",
		}
	}

	client, _ := blizzard.GetClient(ctx, opts["region"], opts["locale"])

	raw, err := client.Get(
		ctx,
		fmt.Sprintf("/data/wow/pvp-season/%d/pvp-leaderboard/%s", seasonId, bracket),
		map[string]string{
			"namespace": "dynamic-" + client.Region,
			"locale":    client.Locale,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("blizzard fetch failed: %w", err)
	}

	var dto response.PvpLeaderboardResponse
	if err := json.Unmarshal(raw, &dto); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return &dto, nil
}
