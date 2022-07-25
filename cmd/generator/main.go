package main

import (
	"fmt"
	"github.com/maxheckel/scare-me-to-sleep/internal/config"
	"github.com/maxheckel/scare-me-to-sleep/internal/db"
	"github.com/maxheckel/scare-me-to-sleep/internal/domain"
	"github.com/maxheckel/scare-me-to-sleep/internal/services"
	"log"
	"time"
)

func main() {
	cfg, err := config.Load()
	database, err := db.Connect(cfg.DBFile)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrate(database)
	if err != nil {
		log.Fatal(err)
	}
	nextPrompt := domain.Prompt{}
	err = database.First(&nextPrompt, "answered = false").Order("priority").Error
	if err != nil {
		panic(err)
	}

	client := services.NewOpenAIClient(cfg)
	for range make([]int, 25) {
		resp, err := client.Generate(&nextPrompt)
		if err != nil {
			panic(err)
		}
		err = database.Create(&resp).Error
		if err != nil {
			panic(err)
		}
		fmt.Println("Generated Answer")
	}
	nextPrompt.Answered = true
	now := time.Now()
	nextPrompt.AnsweredOn = &now
	database.Updates(&nextPrompt)

}
