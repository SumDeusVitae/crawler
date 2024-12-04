package main

import (
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)

	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}

	var links []string

	findAnchorTags(doc, &links)
	var final []string
	for _, link := range links {
		newUrl := link
		if !strings.HasPrefix(link, "http") {
			newUrl = rawBaseURL + link
		}
		final = append(final, newUrl)

	}

	return final, nil

}

func findAnchorTags(n *html.Node, links *[]string) {
	// If the node is an element and it's an <a> tag
	if n.Type == html.ElementNode && n.Data == "a" {
		// Print the href attribute of the <a> tag
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				*links = append(*links, attr.Val)
			}
		}
	}

	// Recur for each child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findAnchorTags(c, links)
	}
}
