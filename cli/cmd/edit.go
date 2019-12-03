package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a cheat",
	Long: `Edit a cheat's "command", "name", and/or "description".
         Pass the "id" of the cheat to be edited.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Edit")
		fmt.Println(args)
	},
}
