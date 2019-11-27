package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	// rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
}

var versionCmd = &cobra.Command{
	Use: "version",
	// TraverseChildren: true, parse local flags on each command before executing the target command.
	// Args: cobra.MinimumNArgs(1),
	Short: "Print the version number of Cheat",
	Long:  `All software has versions. This is Cheat's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1.0")
	},
}
