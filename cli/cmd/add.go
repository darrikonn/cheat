package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"cheat/cli/db"
	"cheat/cli/utils"
)

var addDescription string = strings.TrimSpace(`
Add a new cheat to your cheatsheet. Your cheat should
include the cheat "command", "name", and "description".
`)

var (
	addFlags = &struct {
		name        string
		description string
		weight      int
	}{}
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addFlags.name, "name", "n", "", "name of the cheat")
	addCmd.Flags().StringVarP(&addFlags.description, "description", "d", "", "description of the cheat")
	addCmd.Flags().IntVarP(&addFlags.weight, "weight", "w", 0, "weight of the cheat; used for sorting query results")
}

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a new cheat",
	Long:    addDescription,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := db.AddCheat(args[0], addFlags.name, addFlags.description, addFlags.weight)
		utils.Render(
			"Created cheat {BOLD}{GREEN}{id}{RESET}",
			map[string]string{"id": id},
		)
	},
}
