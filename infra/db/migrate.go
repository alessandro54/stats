package db

import (
	"github.com/alessandro54/stats/infra/db/migrations"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"log"
)

func DropDatabase(db *gorm.DB) {
	if err := db.Migrator().DropTable("pvp_leaderboard_snapshots"); err != nil {
		log.Fatalf("❌ Could not drop database: %v", err)
	}

	log.Println("✅ Database dropped successfully")
}

func RunMigrations(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migrations.CreateLeaderboardSnapshotsMigration(),
		// Add future migrations here in order
	})

	if err := m.Migrate(); err != nil {
		log.Fatalf("❌ Could not run migrations: %v", err)
	}

	log.Println("✅ Database migrated successfully")
}
