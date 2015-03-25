package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := loadConfig()
	handleErr(err)

	entries, err := getEntries(config)
	handleErr(err)

	fmt.Printf("%v / %v", sumHours(entries), config.ExpectedDailyHours)
}

func sumHours(entries []Entry) float32 {
	var total float32

	for _, entry := range entries {
		total += entry.Hours
	}

	return total
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		os.Exit(1)
	}
}
