package model

import "github.com/alessandro54/stats/internal/common/model"

type Character struct {
	BlizzardID uint   `json:"blizzard_id" gorm:"uniqueIndex:idx_blizzard_region;not null"`
	Name       string `json:"name" gorm:"not null"`
	RealmSlug  string `json:"realm_slug" gorm:"not null"`
	RealmID    uint   `json:"realm_id" gorm:"not null"`
	Region     string `json:"region" gorm:"uniqueIndex:idx_blizzard_region;not null"`

	model.BaseModel
}
