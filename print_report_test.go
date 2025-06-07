package main

import (
	"reflect"
	"testing"
)

func TestPrintReport(t *testing.T) {
	tests := []struct {
		name     string
		pages    map[string]int
		expected []Pages
	}{
		{
			name: "page numbers",
			pages: map[string]int{"www.firstpage.com": 10, "www.secondpage.com": 7,
				"www.fifthpage.com": 3, "www.fourthpage.com": 4,
				"www.thirdpage.com": 6},
			expected: []Pages{
				{Host: "www.firstpage.com", Visits: 10},
				{Host: "www.secondpage.com", Visits: 7},
				{Host: "www.thirdpage.com", Visits: 6},
				{Host: "www.fourthpage.com", Visits: 4},
				{Host: "www.fifthpage.com", Visits: 3},
			},
		},
		{
			name: "all equal",
			pages: map[string]int{"www.firstpage.com": 1, "www.secondpage.com": 1,
				"www.fifthpage.com": 1, "www.fourthpage.com": 1,
				"www.thirdpage.com": 1},
			expected: []Pages{
				{Host: "www.firstpage.com", Visits: 1},
				{Host: "www.secondpage.com", Visits: 1},
				{Host: "www.fifthpage.com", Visits: 1},
				{Host: "www.fourthpage.com", Visits: 1},
				{Host: "www.thirdpage.com", Visits: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortPages(tt.pages)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Expected %+v, but got %+v", tt.expected, got)
			}
		})
	}

}
