package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

func crawlWeb(ctx context.Context, urlChan <-chan string) <-chan *string {
	results := make(chan *string)
	var wg sync.WaitGroup
	for url := range urlChan {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			res, err := http.Get(url)
			if err != nil {
				log.Println(err)
				return
			}
			content, err := io.ReadAll(res.Body)
			if err != nil {
				log.Println(err)
				return
			}
			res.Body.Close()
			contentStr := string(content)
			select {
			case results <- &contentStr:
			case <-ctx.Done():
				return
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func main() {
	urls := []string{"https://www.google.com/", "https://yandex.ru"}
	ctx, cancel := context.WithCancel(context.Background())
	urlChan := make(chan string, 8)

	go func() {
		defer close(urlChan)
		for _, url := range urls {
			urlChan <- url
		}
	}()
	results := crawlWeb(ctx, urlChan)

	for result := range results {
		fmt.Println(*result)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	select {
	case <-c:
		cancel()
	case <-results:
	}
}
