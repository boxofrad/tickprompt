package main

import "os"

const LOCK_FILE = CACHE_FILE + ".lock"

func update() {
	if updateInProgress() {
		return
	}

	createLockFile()

	config, err := LoadConfig()
	handleErr(err)

	entries, err := GetEntries(config)
	handleErr(err)

	hours := sumHours(entries)
	err = NewCache(hours).WriteToFile()
	handleErr(err)

	deleteLockFile()
}

func sumHours(entries []Entry) float32 {
	var total float32

	for _, entry := range entries {
		total += entry.Hours
	}

	return total
}

func createLockFile() {
	path, err := FilePath(LOCK_FILE)
	handleErr(err)

	file, err := os.Create(path)
	handleErr(err)

	file.Close()
}

func deleteLockFile() {
	path, err := FilePath(LOCK_FILE)
	handleErr(err)

	os.Remove(path)
}

func updateInProgress() bool {
	path, err := FilePath(LOCK_FILE)
	handleErr(err)

	_, err = os.Stat(path)
	return !os.IsNotExist(err)
}
