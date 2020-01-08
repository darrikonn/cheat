package utils

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/viper"

	"cheat/cli/exceptions"
)

// GetEnvWithFallback : gets environment variable with a fallback
func GetEnvWithFallback(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetUserInputFromEditor : creates a temp file where user can edit
// their text in their preferred editor, and returns the
// output from that temp file
func GetUserInputFromEditor(existingContent string) string {
	// Create tmp file
	file, err := ioutil.TempFile("/tmp", "cheat")
	if err != nil {
		panic(exceptions.CheatException("Could not create a tmp file at \"/tmp/…\"", err))
	}
	defer func() {
		err = os.Remove(file.Name())
		if err != nil {
			panic(exceptions.CheatException("Could not delete tmp file: \""+file.Name()+"\"", err))
		}
	}()

	// Write existing value to tmp file
	err = ioutil.WriteFile(file.Name(), []byte(existingContent), 0644)
	if err != nil {
		panic(exceptions.CheatException("Could not write to tmp file: \""+file.Name()+"\"", err))
	}

	// Open file for editing
	cmd := exec.Command(viper.GetString("editor"), file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	buf, err := ioutil.ReadFile(file.Name())
	if err != nil {
		panic(exceptions.CheatException("Could not read from tmp file: \""+file.Name()+"\"", err))
	}
	return string(buf)
}

// Check : handle errors for defer functions
func Check(f func() error, reason string) {
	if err := f(); err != nil {
		panic(exceptions.CheatException(reason, err))
	}
}

// ContainsAny : returns true if target contains any of strings
func ContainsAny(target []string, matches ...string) bool {
	for _, t := range target {
		for _, match := range matches {
			if t == match {
				return true
			}
		}
	}
	return false
}
