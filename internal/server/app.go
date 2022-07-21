package server

import (
	"github.com/maxheckel/scare-me-to-sleep/internal/config"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Config *config.Config
}
