package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func configure(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing base URL: %v", err)
	}
	return &config{
		pages:              make(map[string]int),
		baseURL:            base,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}, nil
}
