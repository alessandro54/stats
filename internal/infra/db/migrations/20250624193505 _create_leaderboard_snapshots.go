package migrations

import (
	"github.com/alessandro54/stats/internal/playervsplayer/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreatePvpLeaderboardSnapshotsMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250624193505_create_leaderboard_snapshots",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.PvpLeaderboardSnapshot{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("pvp_leaderboard_snapshots")
		},
	}
}
