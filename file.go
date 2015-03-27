package main

import (
	"os/user"
	"path"
)

func FilePath(name string) (string, error) {
	user, err := user.Current()

	if err != nil {
		return "", err
	}

	return path.Join(user.HomeDir, name), nil
}
