package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var username string
	flag.StringVar(&username, "username", "", "input your username")

	flag.Usage = func() {
		_, err := fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		if err != nil {
			return
		}
		fmt.Println("  -username string\n    input your username (required)")
	}

	// Parse the flags
	flag.Parse()

	if username == "" {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Username:", username)
}
