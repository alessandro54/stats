package model

import "github.com/alessandro54/stats/internal/common/model"

type Specialization struct {
	model.BaseModel

	CharacterID uint `gorm:"not null;foreignKey:CharacterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"character_id"`
	Name        string
	Class       string
}
