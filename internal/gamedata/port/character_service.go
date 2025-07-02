package port

import (
	"context"
	"github.com/alessandro54/stats/internal/gamedata/model"
)

type CharacterService interface {
	GetOrFetch(ctx context.Context, character *model.Character) (*model.Character, error)
}
