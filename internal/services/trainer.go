package services

import (
	"encoding/json"
	"fmt"
	"github.com/maxheckel/scare-me-to-sleep/internal/domain"
	"io/fs"
	"log"
	"path/filepath"
)

func CreateFile() error {
	lines := []string{}
	err := filepath.Walk("data/threads", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("File Name: %s\n", info.Name())
		data, err := RetrieveFile("data/threads/" + info.Name())
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		prompt := data[0]
		repliesRaw := data[1]
		replies := domain.Thread{}.Data.Children
		err = json.Unmarshal([]byte(repliesRaw), &replies)
		if err != nil {
			log.Fatalf(err.Error())
		}
		for _, reply := range replies {
			if reply.Data.Body == "" || prompt == "" {
				continue
			}
			line := domain.FineTune{
				Prompt:     prompt,
				Completion: reply.Data.Body,
			}
			lineStr, err := json.Marshal(line)
			if err != nil {
				return err
			}
			lines = append(lines, string(lineStr))
		}
		fmt.Println("Added " + info.Name())
		return nil
	})
	if err != nil {
		return err
	}
	StoreFile(lines, "data/training")
	return nil
}
