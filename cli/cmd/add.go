package cmd

import (
	"github.com/spf13/cobra"

	"cheat/cli/db"
	"cheat/cli/utils"
)

// flags
var (
	name        string
	description string
	weight      int
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&name, "name", "n", "", "name of the cheat")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "description of the cheat")
	addCmd.Flags().IntVarP(&weight, "weight", "w", 0, "weight of the cheat; used for sorting query results")
}

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a new cheat",
	Long: `
    Add a new cheat to your cheatsheet. Your cheat should
    include the cheat "command", "name", and "description".
  `,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := db.AddCheat(args[0], name, description, weight)
		utils.Render(
			"Created cheat {BOLD}{GREEN}{id}{RESET}",
			map[string]string{"id": id},
		)
	},
}
