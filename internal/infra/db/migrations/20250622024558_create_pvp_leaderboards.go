package migrations

import (
	"github.com/alessandro54/stats/internal/gamedata/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreatePVPLeaderboardsMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250622024558_create_pvp_leaderboards",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.PVPLeaderboard{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("pvp_leaderboards")
		},
	}
}
