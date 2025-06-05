package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	header := resp.Header[http.CanonicalHeaderKey("content-type")]
	if resp.StatusCode >= 399 {
		return "", fmt.Errorf("error getting url")
	}

	if len(header) == 0 || !strings.Contains(header[0], "text/html") {
		return "", fmt.Errorf("missing or invalid Content-Type")
	}

	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("unable to read body")
	}

	return string(html), nil
}
