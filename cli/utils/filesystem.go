package utils

import (
	"os"
	"path/filepath"
	"strings"

	"cheat/cli/cmd/exceptions"
)

// ResolvePath : resolves path; relative/expanding
func ResolvePath(path string) string {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		panic(exceptions.CheatException("Could not find home directory", err))
	}

	absPath, err := filepath.Abs(
		strings.ReplaceAll(path, "~", homeDirectory),
	)
	if err != nil {
		panic(exceptions.CheatException("Could not determine path from \""+path+"\"", err))
	}

	return absPath
}

// CreateFile : creates a file in path
func CreateFile(file string) {
	_, err := os.Create(file)
	if err != nil {
		panic(exceptions.CheatException("Could not create file: \""+file+"\"", err))
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
