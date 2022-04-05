package util

import (
	"os"
)

func IsDir(rootPath string) bool {
	if file, err := os.Stat(rootPath); err != nil {
		return false
	} else {
		return file.IsDir()
	}
}

func HomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "/"
	}
	return homeDir
}
