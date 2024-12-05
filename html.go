package main

import (
	"fmt"
	"io"
	"strings"
)

func getHTML(rawURL string, cfg *config) (string, error) {
	if cfg.isCrawlingStopped() {
		return "", fmt.Errorf("crawling stopped, aborting request to %s", rawURL)
	}
	resp, err := cfg.client.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()
	// checking status
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("error status code: %v\n", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("expected Content-Type 'text/html', but got '%s'", contentType)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading body: %v", err)
	}

	return string(body), nil
}
