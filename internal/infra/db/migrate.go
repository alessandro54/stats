package db

import (
	"github.com/alessandro54/stats/internal/infra/db/migrations"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"log"
)

func migrationsList() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		migrations.CreateLeaderboardSnapshotsMigration(),
		migrations.CreatePVPSeasonsMigration(),
		migrations.CreatePVPLeaderboardsMigration(),
		// Add more migrations here
	}
}

func NewMigrator(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, migrationsList())
}

func RunMigrations(db *gorm.DB) {
	m := NewMigrator(db)
	if err := m.Migrate(); err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("✅ Database migrated successfully")
}
