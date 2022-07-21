package config

import (
	"errors"
	"os"
)

type Config struct {
	DBFile    string
	OpenAIKey string
}

func Load() (*Config, error) {
	dbFile, found := os.LookupEnv("DB_NAME")
	if !found {
		return nil, errors.New("could not find database filename, set ENV variable DB_NAME to be equal to the path to a sqlite db to run app.")
	}

	openAIKey, found := os.LookupEnv("OPEN_AI_KEY")
	if !found {
		return nil, errors.New("no OPEN_AI_KEY configured")
	}

	return &Config{
		DBFile:    dbFile,
		OpenAIKey: openAIKey,
	}, nil
}