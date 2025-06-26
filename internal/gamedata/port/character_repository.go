package port

import (
	"context"
	"github.com/alessandro54/stats/internal/gamedata/model"
)

type CharacterRepository interface {
	FindOneByID(ctx context.Context, characterID uint) (*model.Character, error)
	FindOneByBlizzardID(ctx context.Context, blizzardID string) (*model.Character, error)
	FindOrInsert(ctx context.Context, character *model.Character) (*model.Character, error)

	Save(ctx context.Context, character *model.Character) error

	Delete(ctx context.Context, characterID uint) error
}
