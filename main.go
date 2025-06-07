package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//argsWithProPath := os.Args
	cmd := os.Args[1:]

	if len(cmd) < 3 {
		fmt.Println("missing all arguments")
		os.Exit(1)
	} else if len(cmd) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else if len(cmd) == 3 {
		fmt.Printf("starting crawl of: %v\n", cmd[0])
		maxConcurrency, err := strconv.Atoi(cmd[1])
		if err != nil {
			fmt.Printf("error retriving max concurrency")
			return
		}
		maxPages, err := strconv.Atoi(cmd[2])
		if err != nil {
			fmt.Printf("error retrieving page max")
		}
		cfg, err := configure(cmd[0], maxConcurrency, maxPages)
		if err != nil {
			fmt.Printf("Error - configure: %v", err)
			return
		}
		cfg.wg.Add(1)
		go cfg.crawlPage(cmd[0])
		cfg.wg.Wait()

		cfg.printReport(cfg.pages, cfg.baseURL.String())

	}

}
