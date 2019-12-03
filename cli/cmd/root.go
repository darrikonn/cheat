package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"cheat/cli/cmd/exceptions"
	db "cheat/cli/db"
	"cheat/cli/utils"
)

var (
	// Used for flags.
	cfgFile string
	verbose bool

	rootCmd = &cobra.Command{
		Use:     "cheat",
		Version: "0.1.0",
		Short:   "Cheat is a personal cheatsheet manager",
		Long: `A fast and flexible cheatsheet manager build with
           Go. Complete documentation is available at
           https://github.com/darrikonn/cheat/api.md`,
	}
)

func errorHandling() {
	err := recover()
	if err != nil {
		switch err.(type) {
		case *exceptions.AbortType:
			fmt.Println("Abort!")
		default:
			if verbose {
				panic(err)
			} else {
				utils.Render("{RED}Error{RESET}:", nil)
				fmt.Println(err)
				os.Exit(1)
			}
		}
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cheat.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose output")
	//
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.BindPFlag("useVerbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := homedir.Dir()
		if err != nil {
			// er(err)
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cheat" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cheat")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
