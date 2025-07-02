package service

import (
	"context"
	"fmt"
	pvpseason "github.com/alessandro54/stats/internal/dataextraction/adapter/blizzard/gamedata"
	"github.com/alessandro54/stats/internal/dataextraction/adapter/blizzard/gamedata/response"
	characterModel "github.com/alessandro54/stats/internal/gamedata/model"
	characterPort "github.com/alessandro54/stats/internal/gamedata/port"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
	"github.com/alessandro54/stats/internal/playervsplayer/port"
)

type pvpLeaderboardServiceImpl struct {
	repo         port.PvpLeaderboardRepository
	characterSvc characterPort.CharacterService
}

func (p *pvpLeaderboardServiceImpl) GetOrFetch(ctx context.Context, seasonID uint, bracket string, region string) (*model.PvpLeaderboard, error) {
	leaderboard, err := p.repo.FindBySeasonAndBracket(ctx, seasonID, bracket, region)
	if err != nil {
		return nil, err
	}
	if leaderboard != nil {
		return leaderboard, nil
	}

	dto, err := pvpseason.FetchLeaderboard(ctx, seasonID, bracket, map[string]string{
		"region": region,
		"locale": "en_US",
	})

	if err != nil {
		return nil, fmt.Errorf("blizzard fetch failed: %w", err)
	}

	newLeaderboard, err := p.buildLeaderboardFromDTO(ctx, dto, region)
	if err != nil {
		return nil, err
	}

	if err := p.repo.Insert(ctx, newLeaderboard); err != nil {
		return nil, fmt.Errorf("failed to insert leaderboard: %w", err)
	}

	return newLeaderboard, nil
}

func (p *pvpLeaderboardServiceImpl) buildLeaderboardFromDTO(
	ctx context.Context,
	dto *response.PvpLeaderboardResponse,
	region string,
) (*model.PvpLeaderboard, error) {

	leaderboard := &model.PvpLeaderboard{
		PvpSeasonID: dto.Season.ID,
		Bracket:     dto.Name,
		Region:      region,
	}

	for _, entry := range dto.Entries {
		println(entry.Character.Name)
		char := &characterModel.Character{
			BlizzardID: entry.Character.ID,
			Name:       entry.Character.Name,
			RealmSlug:  entry.Character.Realm.Slug,
			RealmID:    entry.Character.Realm.ID,
			Region:     region,
		}

		storedChar, err := p.characterSvc.GetOrFetch(ctx, char)
		if err != nil {
			continue
		}

		leaderboard.Entries = append(leaderboard.Entries, model.PvpLeaderboardEntry{
			CharacterID: storedChar.ID,
			Rating:      entry.Rating,
			Rank:        entry.Rank,
			Played:      entry.SeasonMatchStatistics.Played,
			Won:         entry.SeasonMatchStatistics.Won,
			Lost:        entry.SeasonMatchStatistics.Lost,
		})
	}

	fmt.Println("Entries to insert:", len(leaderboard.Entries))

	return leaderboard, nil
}

func NewPvpLeaderboardService(repo port.PvpLeaderboardRepository, svc characterPort.CharacterService) port.PvpLeaderboardService {
	return &pvpLeaderboardServiceImpl{
		repo:         repo,
		characterSvc: svc,
	}
}
