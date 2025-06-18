package migrations

import (
	"github.com/alessandro54/stats/internal/gameinfo/domain/entity"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateLeaderboardSnapshotsMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250617203629_create_leaderboard_snapshots",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&entity.LeaderboardSnapshot{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("pvp_leaderboard_snapshots")
		},
	}
}
