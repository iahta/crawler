package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
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
	for n := range htmlNodes.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					u, err := url.Parse(a.Val)
					if err != nil {
						return nil, err
					}
					base, err := url.Parse(rawBaseURL)
					if err != nil {
						return nil, err
					}
					urlBase := base.ResolveReference(u)
					urls = append(urls, urlBase.String())
				}
			}
		}
	}
	return urls, nil
}
