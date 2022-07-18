package main

import (
	"fmt"
	"github.com/maxheckel/scare-me-to-sleep/internal/services"
)

func main() {
	for _, t := range services.FindThreads() {
		fmt.Printf("Crawling %s\n", t)
		err := services.CrawlThread(t)
		if err != nil {
			fmt.Println(err.Error())
		}

	}

}
