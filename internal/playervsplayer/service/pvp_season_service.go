package service

import (
	"context"
	"encoding/json"
	"fmt"
	pvpseason "github.com/alessandro54/stats/internal/dataextraction/adapter/blizzard/gamedata"
	"github.com/alessandro54/stats/internal/playervsplayer/model"
	"github.com/alessandro54/stats/internal/playervsplayer/port"
	"time"
)

type PvpSeasonDTO struct {
	ID                uint   `json:"id"`
	SeasonStartUnixMs int64  `json:"season_start_timestamp"`
	SeasonEndUnixMs   int64  `json:"season_end_timestamp"`
	SeasonName        string `json:"season_name"`
}

type pvpSeasonServiceImpl struct {
	repo port.PvpSeasonRepository
}

func (p *pvpSeasonServiceImpl) FetchOrInsert(ctx context.Context, blizzardID uint) (*model.PvpSeason, error) {
	if season, err := p.findExisting(ctx, blizzardID); season != nil || err != nil {
		return season, err
	}

	data, err := pvpseason.FetchPvpSeason(ctx, blizzardID, "us")
	if err != nil {
		return nil, fmt.Errorf("blizzard fetch failed: %w", err)
	}

	var dto PvpSeasonDTO
	if err := json.Unmarshal(data, &dto); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	start := time.UnixMilli(dto.SeasonStartUnixMs).UTC()

	var end *time.Time
	if dto.SeasonEndUnixMs > 0 {
		t := time.UnixMilli(dto.SeasonEndUnixMs).UTC()
		end = &t
	}

	newSeason := &model.PvpSeason{
		BlizzardID: dto.ID,
		Name:       dto.SeasonName,
		StartTime:  start,
		EndTime:    end,
	}

	if err := p.repo.Insert(ctx, newSeason); err != nil {
		return nil, err
	}
	return newSeason, nil
}

func (p *pvpSeasonServiceImpl) findExisting(ctx context.Context, id uint) (*model.PvpSeason, error) {
	season, err := p.repo.FindByBlizzardID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find season: %w", err)
	}
	return season, nil
}

func NewPvpSeasonService(repo port.PvpSeasonRepository) port.PvpSeasonService {
	return &pvpSeasonServiceImpl{
		repo: repo,
	}
}
