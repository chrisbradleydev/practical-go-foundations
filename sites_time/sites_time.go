package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// siteTime("https://google.com")

	urls := []string{
		"https://facebook.com",
		"https://amazon.com",
		"https://apple.com",
		"https://netflix.com",
		"https://google.com",
	}

	wg.Add(len(urls))
	for _, url := range urls {
		// wg.Add(1)
		url := url
		go func() {
			defer wg.Done()
			siteTime(url)
		}()
	}
	wg.Wait()

	// another option if you only care about errors
	// https://pkg.go.dev/golang.org/x/sync/errgroup
}

func siteTime(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR: %s -> %s", url, err)
	}
	defer resp.Body.Close()
	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Printf("ERROR: %s -> %s", url, err)
	}

	duration := time.Since(start)
	log.Printf("INFO: %s -> %s", url, duration)
}