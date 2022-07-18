package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/maxheckel/scare-me-to-sleep/internal/domain"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func CrawlThread(uri string) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.reddit.com%s.json", uri), nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Access-Control-Allow-Origin", "*")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	respArr := []domain.Thread{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &respArr)
	if err != nil {
		return err
	}
	if len(respArr) < 2 {
		return errors.New("expected thread and replies but only got thread")
	}
	thread := respArr[0]
	replies := respArr[1]
	if thread.Data.Children[0].Data.Title == "[deleted by user]" {
		fmt.Println("Original prompt missing, skipping")
		return nil
	}
	fileArr := []string{}
	fileArr = append(fileArr, thread.Data.Children[0].Data.Title)
	actualChildren := domain.Thread{}

	for _, r := range replies.Data.Children {
		if strings.Contains(r.Data.Body, "**Attention! [Serious] Tag Notice**") ||
			r.Data.Author == "AutoModerator" ||
			r.Data.Author == "[deleted]" ||
			r.Data.Body == "[removed]" {
			continue
		}
		actualChildren.Data.Children = append(actualChildren.Data.Children, r)
	}
	jsonString, err := json.Marshal(actualChildren.Data.Children)
	if err != nil {
		return err
	}
	fileArr = append(fileArr, string(jsonString))
	err = StoreFile(fileArr, fmt.Sprintf("data/threads/%s", thread.Data.Children[0].Data.ID))
	if err != nil {
		return err
	}
	return nil
}

func FindThreads() []string {
	cached, err := RetrieveFile("data/threads_list")
	if err == nil {
		return cached
	}
	currentURL := "https://www.reddit.com/r/creepyaskreddit.json"
	depth := 118
	var iterationThreads []string
	threads := map[string]bool{}
TOP:
	for i := 0; i < depth; i++ {
		fmt.Printf("Crawling %s\n", currentURL)
		iterationThreads, currentURL, err = crawlThreadList(currentURL)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		okCount := 0
		for _, thread := range iterationThreads {
			// As soon as we've already seen a thread before that means were' at the beginning again
			if _, ok := threads[thread]; ok {
				okCount++
			}
			if okCount > 5 {
				fmt.Println("Already been here, max threads_list reached!")
				break TOP
			}
			threads[thread] = true
		}
		time.Sleep(1 * time.Second)
	}
	keys := make([]string, len(threads))

	i := 0
	for k := range threads {
		keys[i] = k
		i++
	}
	StoreFile(keys, "data/threads_list")
	return keys
}

func crawlThreadList(url string) (threads []string, next string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Access-Control-Allow-Origin", "*")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return threads, next, err
	}
	threadsData := domain.Threads{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return threads, next, err
	}
	err = json.Unmarshal(body, &threadsData)
	if err != nil {
		return threads, next, err
	}
	next = fmt.Sprintf("https://www.reddit.com/r/CreepyAskReddit.json?count=25&after=%s", threadsData.Data.After)
	for _, t := range threadsData.Data.Children {
		if len(t.Data.CrosspostParentList) == 0 {
			continue
		}
		threads = append(threads, t.Data.CrosspostParentList[0].Permalink)
	}
	return threads, next, err
}
