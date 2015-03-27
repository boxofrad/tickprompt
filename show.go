package main

import (
	"fmt"
	"os/exec"

	"github.com/kardianos/osext"
)

func show() {
	config, err := LoadConfig()
	silentHandleErr(err)

	cache, err := ReadCacheFromFile()
	silentHandleErr(err)

	if cache.HasExpired(config.CacheTTL) {
		updateInBackground()
	}

	fmt.Print(cache.Hours)
}

func updateInBackground() {
	if !updateInProgress() {
		bin, err := osext.Executable()
		silentHandleErr(err)
		exec.Command(bin, "update").Start()
	}
}
