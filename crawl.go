package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	defer cfg.wg.Done()                         // Always decrement WaitGroup
	defer func() { <-cfg.concurrencyControl }() // Always release concurrency control

	// Early exit if crawling has stopped
	if cfg.isCrawlingStopped() {
		return
	}

	cfg.concurrencyControl <- struct{}{} // Acquire concurrency slot

	// Check if maximum pages limit is reached
	if cfg.pagesLen() >= cfg.maxPages {
		cfg.stopCrawlingNow()
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	// Ensure the URL belongs to the base domain
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - normalizeURL: %v\n", err)
		return
	}

	// Check if the URL has been visited before
	if !cfg.addPageVisit(normalizedURL) {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL, cfg)
	if err != nil {
		fmt.Printf("Error - getHTML: %v\n", err)
		return
	}

	nextURLs, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("Error - getURLsFromHTML: %v\n", err)
		return
	}

	// Spawn goroutines for unvisited URLs
	for _, nextURL := range nextURLs {
		if cfg.isCrawlingStopped() {
			return
		}
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}
