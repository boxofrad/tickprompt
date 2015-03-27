package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		usage()
	}

	if args[0] == "update" {
		update()
	}

	if args[0] == "show" {
		show()
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		os.Exit(1)
	}
}

func silentHandleErr(err error) {
	if err != nil {
		os.Exit(1)
	}
}
