package main

import (
	"fmt"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme and slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme no s",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme no s with slash",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "space",
			inputURL: "   http://blog.boot.dev/path/   ",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "space 2",
			inputURL: "http://  blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "just boot",
			inputURL: "http://blog.boot.dev/",
			expected: "blog.boot.dev",
		},
		{
			name:     "query parameters",
			inputURL: "https://blog.boot.dev/path?param=value",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "url fragments",
			inputURL: "https://blog.boot.dev/path#section",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "url fragments and query",
			inputURL: "https://blog.boot.dev/path?param=value#section",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "default https port",
			inputURL: "https://blog.boot.dev:443/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "default http port",
			inputURL: "http://blog.boot.dev:80/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "non-standard port",
			inputURL: "https://blog.boot.dev:8080/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "no scheme",
			inputURL: "blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "domain only",
			inputURL: "blog.boot.dev",
			expected: "blog.boot.dev",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			fmt.Printf("%v: expected: %v, actual %v", tc.name, tc.expected, actual)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLSFromHTML(tc.inputBody, tc.inputURL)
			fmt.Printf("%v: expected: %v, actual %v", tc.name, tc.expected, actual)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			for i, url := range actual {
				if url != tc.expected[i] {
					t.Errorf("Test %v - %s FAIL: expected URL %v, actual: %v", i, tc.name, tc.expected[i], url)
				}
			}

		})
	}
}
