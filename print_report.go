package main

import (
	"fmt"
	"sort"
)

type Pages struct {
	Host   string
	Visits int
}

func (cfg *config) printReport(pages map[string]int, baseURL string) {
	fmt.Printf("=============================\n")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Printf("=============================\n")
	sortedPages := sortPages(pages)
	for _, page := range sortedPages {
		fmt.Printf("Found %v internal links to %s\n", page.Visits, page.Host)
	}

}

func sortPages(pages map[string]int) []Pages {
	sortPages := make([]Pages, 0, len(pages))

	for page, visit := range pages {
		sortPages = append(sortPages, Pages{page, visit})
	}
	sort.SliceStable(sortPages, func(i, j int) bool {
		if sortPages[i].Visits == sortPages[j].Visits {
			return sortPages[i].Host < sortPages[j].Host
		} else {
			return sortPages[i].Visits > sortPages[j].Visits
		}
	})

	return sortPages
}
