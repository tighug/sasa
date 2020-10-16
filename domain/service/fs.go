package service

import (
	"os"
)

// EnsureDir ...
func EnsureDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

// EnsureFile ...
func EnsureFile(path string) {
	os.OpenFile(path, os.O_RDONLY|os.O_CREATE, os.ModePerm)
}
