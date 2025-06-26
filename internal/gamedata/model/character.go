package model

import "github.com/alessandro54/stats/internal/common/model"

type Character struct {
	BlizzardID string `json:"blizzard_id" gorm:"uniqueIndex;not null"`
	Name       string `json:"name" gorm:"not null"`
	RealmSlug  string `json:"realm_slug" gorm:"not null"`
	RealmID    string `json:"realm_id" gorm:"not null"`
	Region     string `json:"region" gorm:"not null"`

	model.BaseModel
}
