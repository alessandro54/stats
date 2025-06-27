package model

import (
	"github.com/alessandro54/stats/internal/common/model"
)

type PvpLeaderboard struct {
	PvpSeasonID uint       `gorm:"not null;index" json:"season_id"`
	PvpSeason   *PvpSeason `gorm:"foreignKey:PvpSeasonID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Bracket     string     `gorm:"not null"`
	Region      string     `gorm:"not null"`

	PvpLeaderboardSnapshots []PvpLeaderboardSnapshot `gorm:"foreignKey:PvpLeaderboardID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"snapshots"`
	Entries                 []PvpLeaderboardEntry    `gorm:"foreignKey:PvpLeaderboardID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"entries"`

	model.BaseModel
}
