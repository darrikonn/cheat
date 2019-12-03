package main

import (
	"log"
	"os"

	"cheat/cli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("Could not execute commands", err)
		os.Exit(1)
	}
}
