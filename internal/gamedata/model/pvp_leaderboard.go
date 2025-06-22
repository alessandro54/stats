package model

import (
	"gorm.io/gorm"
	"time"
)

type PVPLeaderboard struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	PVPSeasonID uint      `gorm:"uniqueIndex"` // Unique index for season
	PVPSeason   PVPSeason `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Bracket     string    `gorm:"not null"`
	Region      string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null;index"`
}
