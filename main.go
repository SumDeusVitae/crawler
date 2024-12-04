package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		return
	}
	argsLen := len(os.Args)
	maxConcurrency := 5
	maxPages := 10
	if argsLen >= 3 {
		maxConcurrency64, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Printf("%s is not an integer value", os.Args[2])
		} else {
			maxConcurrency = int(maxConcurrency64)
		}
		if argsLen >= 4 {
			maxPages64, err := strconv.ParseInt(os.Args[3], 10, 64)
			if err != nil {
				fmt.Printf("%s is not an integer value", os.Args[3])
			} else {
				maxPages = int(maxPages64)
			}
		}
	}

	url := os.Args[1]

	cfg, err := configure(url, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s\n", url)

	cfg.wg.Add(1)
	go cfg.crawlPage(url)

	// Wait for all HTTP fetches to complete.
	cfg.wg.Wait()

	printReport(cfg.pages, url)

}
