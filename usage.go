package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("usage: tickprompt {update,show}")
	os.Exit(1)
}
