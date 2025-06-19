package service

import (
	"context"
	"encoding/json"
	"github.com/alessandro54/stats/internal/gameinfo/adapter/blizzard"
	"github.com/alessandro54/stats/internal/gameinfo/domain/entity"
	"github.com/alessandro54/stats/internal/gameinfo/domain/port"
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

func (s *snapshotService) FetchFromBlizzardAndSave(ctx context.Context, pvpSeasonId string, pvpBracket string) error {
	blizz := blizzard.GetClient()

	data, err := blizz.FetchPvpLeaderboard(ctx, pvpSeasonId, pvpBracket)

	if err != nil {
		return err
	}

	var result any
	if err := json.Unmarshal(data, &result); err != nil {
		panic("failed to parse Blizzard JSON: " + err.Error())
	}

	print(result)

	snapshot := &entity.LeaderboardSnapshot{
		Mode:      "pvp",
		Bracket:   pvpBracket,
		Data:      data,
		CreatedAt: time.Now(),
	}

	return s.repo.SaveSnapshot(ctx, snapshot)
}
