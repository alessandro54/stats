package migrations

import (
	"github.com/alessandro54/stats/internal/playervsplayer/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreatePvpSeasonsMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250624092902_create_pvp_seasons",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.PvpSeason{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("pvp_seasons")
		},
	}
}
