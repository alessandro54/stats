package model

import (
	"github.com/alessandro54/stats/internal/common/model"
)

type PvpLeaderboard struct {
	PvpSeasonID uint       `gorm:"index:idx_unique_leaderboard,unique;not null" json:"season_id"`
	PvpSeason   *PvpSeason `gorm:"foreignKey:PvpSeasonID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Bracket     string     `gorm:"index:idx_unique_leaderboard,unique;not null" json:"bracket"`
	Region      string     `gorm:"index:idx_unique_leaderboard,unique;not null" json:"region"`

	PvpLeaderboardSnapshots []PvpLeaderboardSnapshot `gorm:"foreignKey:PvpLeaderboardID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"snapshots"`
	Entries                 []PvpLeaderboardEntry    `gorm:"foreignKey:PvpLeaderboardID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"entries"`

	model.BaseModel
}
