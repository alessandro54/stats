package migrations

import (
	"github.com/alessandro54/stats/internal/gamedata/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreatePVPSeasonsMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250622024532_create_pvp_seasons",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.PVPSeason{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("pvp_seasons")
		},
	}
}
