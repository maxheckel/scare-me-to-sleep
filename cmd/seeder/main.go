package main

import (
	"fmt"
	"github.com/maxheckel/scare-me-to-sleep/internal/config"
	"github.com/maxheckel/scare-me-to-sleep/internal/db"
	"github.com/maxheckel/scare-me-to-sleep/internal/domain"
	"github.com/maxheckel/scare-me-to-sleep/internal/services"
	"io/fs"
	"log"
	"path/filepath"
)

func main() {
	lines := []string{}
	filepath.Walk("data/threads", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		data, err := services.RetrieveFile("data/threads/" + info.Name())
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		prompt := data[0]
		lines = append(lines, prompt)
		return nil
	})
	cfg, _ := config.Load()
	database, _ := db.Connect(cfg.DBFile)
	db.Migrate(database)
	for _, promptText := range lines {
		prompt := domain.Prompt{
			Text: promptText,
		}
		err := database.Create(&prompt).Error
		if err != nil {
			fmt.Println(err)
		}
	}
}
