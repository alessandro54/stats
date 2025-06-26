package pvpseason

import (
	"context"
	"github.com/alessandro54/stats/internal/dataextraction/adapter/blizzard"
)

func FetchPvpSeasonIndex(ctx context.Context, opts map[string]string) ([]byte, error) {
	client, _ := blizzard.GetClient(ctx, opts["region"], opts["locale"])
	return client.Get(
		ctx,
		"/data/wow/pvp-season/index",
		map[string]string{
			"namespace": "dynamic-" + client.Region,
			"locale":    client.Locale,
		},
	)
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

func FetchLeaderboard(ctx context.Context, seasonId string, bracket string, opts map[string]string) ([]byte, error) {
	client, _ := blizzard.GetClient(ctx, opts["region"], opts["locale"])
	return client.Get(
		ctx,
		"/data/wow/pvp-season/"+seasonId+"/pvp-leaderboard/"+bracket,
		map[string]string{
			"namespace": "dynamic-" + client.Region,
			"locale":    client.Locale,
		},
	)
}
