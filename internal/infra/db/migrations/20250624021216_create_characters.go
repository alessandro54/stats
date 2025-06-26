package migrations

import (
	"github.com/alessandro54/stats/internal/gamedata/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateCharactersMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250624021216_create_characters",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.Character{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("characters")
		},
	}
}
