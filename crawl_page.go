package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	if cfg.pagesFull() {
		return
	}
	current, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}
	if cfg.baseURL.Hostname() != current.Hostname() {
		return
	}
	normalCurrent, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	if !cfg.addPageVisit(normalCurrent) {
		return
	}

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		return
	}
	//fmt.Printf("%v", html)
	allURLS, err := getURLSFromHTML(html, cfg.baseURL.String())
	if err != nil {
		return
	}

	for _, nextURL := range allURLS {
		fmt.Printf("Crawling: %v\n", nextURL)
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)

	}

}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, visited := cfg.pages[normalizedURL]; visited {
		fmt.Printf("Page Visited: %s\n", normalizedURL)
		cfg.pages[normalizedURL]++
		return false
	}
	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) pagesFull() (full bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages) >= cfg.maxPages
}
