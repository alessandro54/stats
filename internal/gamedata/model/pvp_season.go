package model

import "time"

type PVPSeason struct {
	ID        uint `gorm:"primaryKey"`
	SeasonID  uint `gorm:"uniqueIndex"`
	Name      string
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time
	CreatedAt time.Time `gorm:"not null;index"`
}
