package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/kardianos/osext"
)

func show() {
	config, err := LoadConfig()
	silentHandleErr(err)

	cache, err := ReadCacheFromFile()

	if err == nil {
		if cache.HasExpired(config.CacheTTL) {
			updateInBackground()
		}

		fmt.Print(cache.Hours)
	} else {
		if os.IsNotExist(err) {
			updateInBackground()
		} else {
			silentHandleErr(err)
		}
	}
}

func updateInBackground() {
	if !updateInProgress() {
		bin, err := osext.Executable()
		silentHandleErr(err)
		exec.Command(bin, "update").Start()
	}
}
