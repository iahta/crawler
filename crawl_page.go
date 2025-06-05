package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) error {
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return fmt.Errorf("unable to parse base url\nError: %v", err)
	}
	current, err := url.Parse(rawCurrentURL)
	if err != nil {
		return fmt.Errorf("unable to parse current url\nError: %v", err)
	}
	if base.Hostname() != current.Hostname() {
		return fmt.Errorf("host names do not match")
	}
	normalCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return fmt.Errorf("failed to normalize url\nError: %v", err)
	}
	if _, visited := pages[normalCurrent]; visited {
		pages[normalCurrent]++
		return nil
	}

	pages[normalCurrent] = 1

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		return fmt.Errorf("failed to retreive html\nError: %v", err)
	}
	fmt.Printf("%v", html)
	allURLS, err := getURLSFromHTML(html, rawBaseURL)
	if err != nil {
		return fmt.Errorf("failed to get urls\nError: %v", err)
	}
	for _, nextURL := range allURLS {
		err := crawlPage(rawBaseURL, nextURL, pages)
		if err != nil {
			continue
		}
	}

	return nil
}
