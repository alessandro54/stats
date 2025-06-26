package model

import (
	common "github.com/alessandro54/stats/internal/common/model"
	"github.com/alessandro54/stats/internal/gamedata/model"
)

type PvpLeaderboardEntry struct {
	PvpLeaderboardID uint            `gorm:"not null;index" json:"leaderboard_id"`
	PvpLeaderboard   *PvpLeaderboard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CharacterID uint             `gorm:"not null;index" json:"character_id"`
	Character   *model.Character `gorm:"foreignKey:CharacterID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	Rank   uint `json:"rank"`
	Rating uint `json:"rating"`
	Played uint `json:"played"`
	Won    uint `json:"won"`
	Lost   uint `json:"lost"`

	common.BaseModel
}
