package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(db string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(db), &gorm.Config{})
}
