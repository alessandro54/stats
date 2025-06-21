package entity

import "time"

const (
	ModePvP = "pvp"
	ModePvE = "pve"
)

type LeaderboardSnapshot struct {
	ID        uint      `gorm:"primaryKey"`
	Mode      string    `gorm:"not null;index"`
	Bracket   string    `gorm:"type:varchar(10);not null"`
	Data      []byte    `gorm:"not null"`
	Region    string    `gorm:"type:varchar(5);not null;index"`
	CreatedAt time.Time `gorm:"not null;index"`
}
