package main

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type config struct {
	client             *http.Client
	pages              map[string]int
	baseURL            *url.URL
	maxPages           int
	stopCrawling       bool // New flag to signal stopping
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) isCrawlingStopped() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return cfg.stopCrawling
}

func (cfg *config) stopCrawlingNow() {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if !cfg.stopCrawling {
		fmt.Printf("Maximum pages limit of %d is reached\n", cfg.maxPages)
		cfg.stopCrawling = true
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func configure(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	client := &http.Client{
		Timeout: 3 * time.Second, // Set a 3-second timeout
	}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		client:             client,
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}, nil
}

func (cfg *config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}
