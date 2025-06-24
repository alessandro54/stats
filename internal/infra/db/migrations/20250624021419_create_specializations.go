package migrations

import (
	"github.com/alessandro54/stats/internal/dataaccess/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func CreateSpecializationsMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20250624021419_create_specializations",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&model.Specialization{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("characters")
		},
	}
}
