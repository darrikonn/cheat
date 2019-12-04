package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"

	"cheat/cli/cmd/exceptions"
	db "cheat/cli/db"
	"cheat/cli/utils"
)

var cheatDescription string = strings.TrimSpace(`
A fast and flexible cheatsheet manager build with
Go. Complete documentation is available at
https://github.com/darrikonn/cheat/api.md
`)

var (
	// Used for flags.
	verbose bool

	rootCmd = &cobra.Command{
		Use:     "cheat",
		Version: "0.1.0",
		Short:   "Cheat is a personal cheatsheet manager",
		Long:    cheatDescription,
	}
)

func errorHandling() {
	err := recover()
	if err != nil {
		switch err.(type) {
		case *exceptions.AbortType:
			utils.Render("Abort!", nil)
			msg := err.(*exceptions.AbortType).Error()
			if msg != "" {
				utils.Render("{RED}{msg}{RESET}", map[string]string{"msg": msg})
			}
		default:
			if verbose {
				panic(err)
			} else {
				utils.Render("{RED}Error{RESET}:", nil)
				utils.Render(err.(error).Error(), nil)
			}
		}
		os.Exit(1)
	}
}

// Execute : executes the root command
// that combines all cli subcommands
func Execute() error {
	defer errorHandling()
	defer db.Cleanup()

	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cheat.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose output")
}
