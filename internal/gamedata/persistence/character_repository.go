package persistence

import (
	"context"
	"errors"
	"github.com/alessandro54/stats/internal/gamedata/model"
	"github.com/alessandro54/stats/internal/gamedata/port"
	"gorm.io/gorm"
	"strings"
)

type characterRepositoryImpl struct {
	db *gorm.DB
}

func (c characterRepositoryImpl) Save(ctx context.Context, character *model.Character) error {
	return c.db.WithContext(ctx).Save(character).Error
}

func (c characterRepositoryImpl) FindOrInsert(ctx context.Context, character *model.Character) (*model.Character, error) {
	var existing model.Character

	err := c.db.WithContext(ctx).Where("blizzard_id = ? AND region = ?", character.BlizzardID, character.Region).First(&existing).Error

	if err == nil {
		return &existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err := c.db.WithContext(ctx).Create(character).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return c.FindOneByBlizzardID(ctx, character.BlizzardID)
		}
		return nil, err
	}

	return character, nil
}

func (c characterRepositoryImpl) FindOneByID(ctx context.Context, characterID uint) (*model.Character, error) {
	var character model.Character

	err := c.db.WithContext(ctx).Where("id = ?", characterID).First(&character).Error

	if err != nil {
		return nil, err
	}

	return &character, nil
}

func (c characterRepositoryImpl) FindOneByBlizzardID(ctx context.Context, blizzardID uint) (*model.Character, error) {
	var character model.Character

	err := c.db.WithContext(ctx).Where("blizzardID = ?", blizzardID).First(&character).Error

	if err != nil {
		return nil, err
	}

	return &character, nil
}

func (c characterRepositoryImpl) FindOneByRegionNameRealm(ctx context.Context, region string, name string, realm string) (*model.Character, error) {
	var character model.Character

	err := c.db.WithContext(ctx).Where(
		"region = ? AND name = ? AND realm = ?", region, name, realm,
	).First(&character).Error

	if err != nil {
		return nil, err
	}

	return &character, nil
}

func (c characterRepositoryImpl) Delete(ctx context.Context, characterID uint) error {
	panic("implement me")
}

func NewCharacterRepository(db *gorm.DB) port.CharacterRepository {
	return &characterRepositoryImpl{db: db}
}
