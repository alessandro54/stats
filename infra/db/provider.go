package db

import "gorm.io/gorm"

func ProvideDB() *gorm.DB {
	if DB == nil {
		panic("DB is not initialized. Call db.Connect() before using ProvideDB.")
	}
	return DB
}
