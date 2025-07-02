package repositories

import (
	"context"
	"github.com/alessandro54/stats/internal/gamedata/model"
	"github.com/alessandro54/stats/internal/gamedata/port"
	"gorm.io/gorm"
)

type characterRepositoryImpl struct {
	db *gorm.DB
}

func (c *characterRepositoryImpl) Insert(ctx context.Context, character *model.Character) error {
	return c.db.WithContext(ctx).Create(character).Error
}

func (c *characterRepositoryImpl) FindOneByID(ctx context.Context, characterID uint) (*model.Character, error) {
	var character model.Character

	err := c.db.WithContext(ctx).Where("id = ?", characterID).First(&character).Error

	if err != nil {
		return nil, err
	}

	return &character, nil
}

func (c *characterRepositoryImpl) FindOneByBlizzardID(ctx context.Context, blizzardID uint, region string) (*model.Character, error) {
	var character model.Character

	err := c.db.WithContext(ctx).Where("blizzard_id = ? AND region = ?", blizzardID, region).First(&character).Error

	if err != nil {
		return nil, err
	}

	return &character, nil
}

func (c *characterRepositoryImpl) Delete(ctx context.Context, characterID uint) error {
	panic("implement me")
}

func NewCharacterRepository(db *gorm.DB) port.CharacterRepository {
	return &characterRepositoryImpl{db: db}
}
