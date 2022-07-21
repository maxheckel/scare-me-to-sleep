package main

import (
	"github.com/maxheckel/scare-me-to-sleep/internal/config"
	"github.com/maxheckel/scare-me-to-sleep/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	app, err := server.NewApp(cfg)
	app.Start()
}
