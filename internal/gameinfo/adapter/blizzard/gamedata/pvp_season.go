package pvp_season

import (
	"context"
	"github.com/alessandro54/stats/internal/gameinfo/adapter/blizzard"
)

func FetchPvpSeasonIndex(ctx context.Context) ([]byte, error) {
	client := blizzard.GetClient()
	return client.Get(
		ctx,
		"/data/wow/pvp-season/index",
		map[string]string{
			"namespace": "dynamic-" + client.Region,
			"locale":    "en_US",
		},
	)
}

func FetchLeaderboard(ctx context.Context, seasonId string, bracket string) ([]byte, error) {
	client := blizzard.GetClient()
	return client.Get(
		ctx,
		"/data/wow/pvp-season/"+seasonId+"/pvp-leaderboard/"+bracket,
		map[string]string{
			"namespace": "dynamic-" + client.Region,
			"locale":    "en_US",
		},
	)
}
