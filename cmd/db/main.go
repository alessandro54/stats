// File: cmd/db/main.go
package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/alessandro54/stats/internal/infra/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if len(os.Args) < 2 {
		log.Fatal("Expected command: migrate | rollback")
	}

	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	switch os.Args[1] {
	case "migrate":
		db.RunMigrations(dbConn)
	case "rollback":
		if err := db.NewMigrator(dbConn).RollbackLast(); err != nil {
			log.Fatalf("❌ Rollback failed: %v", err)
		}
		log.Println("⏪ Rollback successful")
	default:
		log.Fatalf("❌ Unknown command: %s", os.Args[1])
	}
}
