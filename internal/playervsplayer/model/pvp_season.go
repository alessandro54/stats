package model

import (
	common "github.com/alessandro54/stats/internal/common/model"
	"time"
)

type PvpSeason struct {
	BlizzardID        uint `gorm:"uniqueIndex" json:"blizzard_id"`
	Name              string
	TotalLeaderboards uint
	StartTime         time.Time  `gorm:"not null"`
	EndTime           *time.Time `gorm:"default:null"`

	common.BaseModel
}
