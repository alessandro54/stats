package db

import (
	"github.com/alessandro54/stats/internal/infra/db/migrations"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"log"
)

func migrationsList() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		migrations.CreateCharactersMigration(),
		migrations.CreatePvpSeasonsMigration(),
		migrations.CreateSpecializationsMigration(),
		migrations.CreatePvpLeaderboardsMigration(),
		migrations.CreatePvpLeaderboardEntriesMigration(),
		migrations.CreatePvpLeaderboardSnapshotsMigration(),
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
