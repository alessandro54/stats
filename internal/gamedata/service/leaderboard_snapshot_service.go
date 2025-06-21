package service

import (
	"context"
	"encoding/json"
	"github.com/alessandro54/stats/internal/gamedata/adapter/blizzard/gamedata"
	"github.com/alessandro54/stats/internal/gamedata/domain/entity"
	"github.com/alessandro54/stats/internal/gamedata/domain/port"
	"time"
)

type snapshotService struct {
	repo port.LeaderboardSnapshotRepository
}

func NewSnapshotService(repo port.LeaderboardSnapshotRepository) port.LeaderboardSnapshotService {
	return &snapshotService{
		repo: repo,
	}
}

func (s *snapshotService) GetAll(ctx context.Context) ([]*entity.LeaderboardSnapshot, error) {
	return s.repo.GetAllSnapshots(ctx)
}

func (s *snapshotService) Save(ctx context.Context, mode string, bracket string, data []byte) error {
	snapshot := &entity.LeaderboardSnapshot{
		Mode:      mode,
		Bracket:   bracket,
		Data:      data,
		CreatedAt: time.Now(),
	}
	return s.repo.SaveSnapshot(ctx, snapshot)
}

func (s *snapshotService) FetchFromBlizzardAndSave(ctx context.Context, pvpSeasonId string, pvpBracket string, region string) error {
	data, err := pvpseason.FetchLeaderboard(ctx, pvpSeasonId, pvpBracket, map[string]string{
		"region": region,
		"locale": "en_US",
	})

	if err != nil {
		return err
	}

	var result any
	if err := json.Unmarshal(data, &result); err != nil {
		panic("failed to parse Blizzard JSON: " + err.Error())
	}

	snapshot := &entity.LeaderboardSnapshot{
		Mode:      "pvp",
		Bracket:   pvpBracket,
		Data:      data,
		Region:    region,
		CreatedAt: time.Now(),
	}

	return s.repo.SaveSnapshot(ctx, snapshot)
}
