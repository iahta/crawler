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
		html, err := getHTML(cmd[0])
		if err != nil {
			fmt.Printf("error getting html\nError: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%v", html)
	}

}
