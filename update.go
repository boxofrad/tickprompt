package main

func update() {
	config, err := LoadConfig()
	handleErr(err)

	entries, err := GetEntries(config)
	handleErr(err)

	hours := sumHours(entries)
	err = NewCache(hours).WriteToFile()
	handleErr(err)
}

func sumHours(entries []Entry) float32 {
	var total float32

	for _, entry := range entries {
		total += entry.Hours
	}

	return total
}
