package model

import "github.com/alessandro54/stats/internal/common/model"

type Specialization struct {
	CharacterID uint       `gorm:"not null;index" json:"character_id"`
	Character   *Character `gorm:"foreignKey:CharacterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	Name  string `gorm:"not null"`
	Class string `gorm:"not null"`

	model.BaseModel
}
