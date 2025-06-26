package model

import (
	"github.com/alessandro54/stats/internal/common/model"
	"time"
)

type PvpSeason struct {
	BlizzardID string `gorm:"uniqueIndex" json:"blizzard_id"`
	Name       string
	StartTime  time.Time
	EndTime    *time.Time

	model.BaseModel
}
