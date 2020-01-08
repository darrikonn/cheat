package utils

import (
	"os"
	"path/filepath"
	"strings"

	"cheat/cli/exceptions"
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

// CreateDir : creates a directory in path
func CreateDir(dir string) {
	err := os.Mkdir(ResolvePath(dir), 0755)
	if err != nil {
		panic(exceptions.CheatException("Could not create directory: \""+dir+"\"", err))
	}
}

// FileExists : check if file exists
func FileExists(name string) bool {
	if _, err := os.Stat(ResolvePath(name)); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// HomeDir : finds the cheat's home directory in order to
// unclutter your $HOME directory
func HomeDir() string {
	var dir string

	// Check if CHEAT_HOME environment is set
	dir = os.Getenv("CHEAT_HOME")
	if dir != "" {
		if !FileExists(dir) {
			panic(
				exceptions.CheatException(
					"CHEAT_HOME environment variable set to \""+dir+"\", but directory does not exist",
					nil,
				),
			)
		}
		return dir
	}

	// Check if XDG_CONFIG_HOME environment is set
	dir = os.Getenv("XDG_CONFIG_HOME")
	if dir != "" {
		if !FileExists(dir) {
			panic(
				exceptions.CheatException(
					"XDG_CONFIG_HOME environment variable set to \""+dir+"\", but directory does not exist",
					nil,
				),
			)
		}
		// Respect the file structure of XDG_CONFIG_HOME
		// and create the "cheat" directory if it doesn't exist
		dir = dir + "/cheat"
		if !FileExists(dir) {
			CreateDir(dir)
		}
		return dir
	}

	// Else default to home directory
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(exceptions.CheatException("Could not find home directory", err))
	}
	return dir
}
