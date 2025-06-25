package model

import (
	"github.com/alessandro54/stats/internal/common/model"
	"time"
)

type PvpSeason struct {
	BlizzardID uint `gorm:"uniqueIndex" json:"blizzard_id"`
	Name       string
	StartTime  time.Time
	EndTime    *time.Time

	model.BaseModel
}
