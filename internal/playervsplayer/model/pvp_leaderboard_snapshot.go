package model

import (
	common "github.com/alessandro54/stats/internal/common/model"
)

type PvpLeaderboardSnapshot struct {
	ID   uint   `gorm:"primaryKey"`
	Data []byte `gorm:"not null"`

	PvpLeaderboardID uint            `gorm:"not null;index" json:"pvp_leaderboard_id"`
	PvpLeaderboard   *PvpLeaderboard `gorm:"foreignKey:PvpLeaderboardID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	common.BaseModel
}
