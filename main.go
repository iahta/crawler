package main

import (
	"fmt"
	"os"
)

func main() {
	//argsWithProPath := os.Args
	cmd := os.Args[1:]
	pages := make(map[string]int)

	if len(cmd) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(cmd) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else if len(cmd) == 1 {
		fmt.Printf("starting crawl of: %v\n", cmd[0])
		err := crawlPage(cmd[0], cmd[0], pages)
		if err != nil {
			fmt.Printf("error crawling html\nError: %v\n", err)
			os.Exit(1)
		}
		for page, count := range pages {
			fmt.Printf("Key: %v, Value: %v\n", page, count)
		}

	}

}
