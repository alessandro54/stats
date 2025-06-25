package migrations

import (
	"github.com/alessandro54/stats/internal/playervsplayer/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreatePvpLeaderboardEntriesMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250624151024_create_pvp_leaderboard_entries",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.PvpLeaderboardEntry{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("pvp_leaderboard_entries")
		},
	}
}
