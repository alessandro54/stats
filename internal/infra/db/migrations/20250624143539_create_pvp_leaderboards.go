package migrations

import (
	"github.com/alessandro54/stats/internal/playervsplayer/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreatePvpLeaderboardsMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250624143539_create_pvp_leaderboards",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.PvpLeaderboard{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("pvp_leaderboards")
		},
	}
}
