package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// ResolvePath : resolves path; relative/expanding
func ResolvePath(path string) string {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	absPath, err := filepath.Abs(
		strings.ReplaceAll(path, "~", homeDirectory),
	)
	if err != nil {
		panic(err)
	}

	return absPath
}

// CreateFile : creates a file in path
func CreateFile(file string) {
	_, err := os.Create(file)
	if err != nil {
		panic(err)
	}
}

// FileExists : check if file exists
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
