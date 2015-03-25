package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := loadConfig()

	if err != nil {
		log.Fatal(err)
	}

	entries, err := getEntries(config)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v / %v", sumHours(entries), config.ExpectedDailyHours)
}

func sumHours(entries []Entry) float32 {
	var total float32

	for _, entry := range entries {
		total += entry.Hours
	}

	return total
}
