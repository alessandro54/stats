package service

import (
	"context"
	"encoding/json"
	"github.com/alessandro54/stats/internal/dataextraction/adapter/blizzard/gamedata"
	port2 "github.com/alessandro54/stats/internal/dataextraction/port"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
)

type snapshotServiceImpl struct {
	repo port2.SnapshotRepository
}

func NewSnapshotService(repo port2.SnapshotRepository) port2.SnapshotService {
	return &snapshotServiceImpl{
		repo: repo,
	}
}

func (s *snapshotServiceImpl) GetAll(ctx context.Context) ([]*model.PvpLeaderboardSnapshot, error) {
	return s.repo.GetAllSnapshots(ctx)
}

func (s *snapshotServiceImpl) GetLatestSeasonByBracket(ctx context.Context, mode string, bracket string) (*model.PvpLeaderboardSnapshot, error) {
	snapshot, err := s.repo.GetLatestSnapshot(ctx, mode, bracket)

	if err != nil {
		return nil, err
	}

	if snapshot == nil {
		return nil, nil // No snapshot found
	}

	return snapshot, nil
}

func (s *snapshotServiceImpl) Save(ctx context.Context, mode string, bracket string, data []byte) error {
	snapshot := &model.PvpLeaderboardSnapshot{
		Data: data,
	}
	return s.repo.SaveSnapshot(ctx, snapshot)
}

func (s *snapshotServiceImpl) FetchFromBlizzardAndSave(ctx context.Context, pvpSeasonId string, pvpBracket string, region string) error {
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

	snapshot := &model.PvpLeaderboardSnapshot{
		Data: data,
	}

	return s.repo.SaveSnapshot(ctx, snapshot)
}
