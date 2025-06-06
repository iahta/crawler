package main

import (
	"fmt"
	"os"
)

func main() {
	//argsWithProPath := os.Args
	cmd := os.Args[1:]

	if len(cmd) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(cmd) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else if len(cmd) == 1 {
		fmt.Printf("starting crawl of: %v\n", cmd[0])
		const maxConcurrency = 3
		cfg, err := configure(cmd[0], maxConcurrency)
		if err != nil {
			fmt.Printf("Error - configure: %v", err)
			return
		}
		cfg.wg.Add(1)
		go cfg.crawlPage(cmd[0])
		cfg.wg.Wait()

		for page, count := range cfg.pages {
			fmt.Printf("Key: %v, Value: %v\n", page, count)
		}

	}

}
