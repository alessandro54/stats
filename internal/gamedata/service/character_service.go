package service

import (
	"context"
	"errors"
	"github.com/alessandro54/stats/internal/gamedata/model"
	"github.com/alessandro54/stats/internal/gamedata/port"
	"gorm.io/gorm"
	"log"
)

type characterServiceImpl struct {
	repo port.CharacterRepository
}

func (c *characterServiceImpl) GetOrFetch(ctx context.Context, character *model.Character) (*model.Character, error) {
	existing, err := c.repo.FindOneByBlizzardID(ctx, character.BlizzardID, character.Region)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return existing, nil
	}

	if err := c.repo.Insert(ctx, character); err != nil {
		log.Printf("❌ Failed to insert character %s: %v", character.Name, err)

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.repo.FindOneByBlizzardID(ctx, character.BlizzardID, character.Region)
		}
		return nil, err
	}

	return character, nil
}

func NewCharacterService(repo port.CharacterRepository) port.CharacterService {
	return &characterServiceImpl{
		repo: repo,
	}
}
