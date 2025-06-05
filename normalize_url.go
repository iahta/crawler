package main

import (
	"fmt"
	"net/url"
	"slices"
	"strings"

	"golang.org/x/net/html"
)

func normalizeURL(urlString string) (string, error) {
	lower := strings.ToLower(urlString)
	cleaned := strings.ReplaceAll(lower, " ", "")
	parsed, err := url.Parse(cleaned)
	if err != nil {
		return "", err
	}
	host := parsed.Hostname()
	if parsed.Path == "/" {
		return fmt.Sprintf("%v", host), nil
	} else {
		trim := strings.TrimSuffix(parsed.Path, "/")
		return fmt.Sprintf("%v%v", host, trim), nil
	}
}

func getURLSFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	htmlReader := strings.NewReader(htmlBody)
	htmlNodes, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}
	urls := []string{}
	extractedUrls, err := extractURLs(htmlNodes, rawBaseURL, urls)
	if err != nil {
		return nil, err
	}
	return extractedUrls, nil
}

func extractURLs(n *html.Node, baseURL string, urlSlice []string) ([]string, error) {
	if n == nil {
		return urlSlice, nil
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, nAtt := range n.Attr {
			if nAtt.Key == "href" {
				if nAtt.Val == "" {
					continue
				}
				u, err := url.ParseRequestURI(nAtt.Val)
				if err != nil {
					continue
				}
				base, err := url.Parse(baseURL)
				if err != nil {
					return nil, err
				}
				urlBase := base.ResolveReference(u)

				if !slices.Contains(urlSlice, urlBase.String()) && verifyURL(urlBase) {
					urlSlice = append(urlSlice, urlBase.String())
				}
			}

		}
	}
	urlSlice, err := extractURLs(n.FirstChild, baseURL, urlSlice)
	if err != nil {
		return nil, err
	}
	urlSlice, err = extractURLs(n.NextSibling, baseURL, urlSlice)
	if err != nil {
		return nil, err
	}
	return urlSlice, nil
}

// possible to send ping? /safe?
func verifyURL(resolvedURL *url.URL) bool {
	scheme := resolvedURL.Scheme
	host := resolvedURL.Host
	if (scheme == "http" || scheme == "https") && host != "" {
		return true
	}
	return false
}
