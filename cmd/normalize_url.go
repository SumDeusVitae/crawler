package main

import "strings"

func normalizeURL(inputUrl string) (string, error) {
	if strings.HasPrefix(inputUrl, "https://") {
		inputUrl = inputUrl[len("https://"):]
	} else if strings.HasPrefix(inputUrl, "http://") {
		inputUrl = inputUrl[len("http://"):]
	}

	if len(inputUrl) > 0 && inputUrl[len(inputUrl)-1] == '/' {
		inputUrl = inputUrl[:len(inputUrl)-1]
	}

	return inputUrl, nil
}
