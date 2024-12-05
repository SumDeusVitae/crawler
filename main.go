package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		fmt.Println("By default maxConcurrency = 5 and maxPages=10")
		return
	}

	url := os.Args[1]
	maxConcurrency := 5
	maxPages := 10

	if len(os.Args) >= 3 {
		if val, err := strconv.Atoi(os.Args[2]); err == nil {
			maxConcurrency = val
		}
	}

	if len(os.Args) >= 4 {
		if val, err := strconv.Atoi(os.Args[3]); err == nil {
			maxPages = val
		}
	}

	cfg, err := configure(url, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v\n", err)
		return
	}

	fmt.Printf("Starting crawl of: %s\n", url)

	cfg.wg.Add(1)
	go cfg.crawlPage(url)

	// Wait for all goroutines to finish
	cfg.wg.Wait()

	printReport(cfg.pages, url)
}
