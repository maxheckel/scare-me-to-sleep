package db

import (
	"github.com/maxheckel/scare-me-to-sleep/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Prompt{}, &domain.Response{})
}
