package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"cheat/cli/db"
	"cheat/cli/utils"
)

var (
	editFlags = &struct {
		ignoreCase bool
		name       string
		weight     int
	}{}
)

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().BoolVarP(&editFlags.ignoreCase, "ignore-case", "i", false, "Case insensitive search")
	editCmd.Flags().StringVarP(&editFlags.name, "name", "n", "", "Rename the cheat")
	editCmd.Flags().IntVarP(&editFlags.weight, "weight", "w", 0, "Weight of the cheat; used for sorting query results")
	editCmd.SetUsageTemplate(createUsageTemplate("cheat [regex] edit [flags]"))
}

var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e"},
	Short:   "Edit a cheat",
	Long: strings.TrimSpace(`
Edit a cheat's "name" and/or "weight". You'll also be
prompted for the cheat's "description" in your preferred editor.
`),

	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cheat := db.GetCheatByName(args[0], editFlags.ignoreCase)

		// Let's first try to rename the cheat
		if cmd.Flags().Changed("name") {
			cheat = db.RenameCheat(cheat.Name, editFlags.name)
		}

		// Then edit description and weight
		db.EditCheat(
			cheat.Name,
			utils.GetUserInputFromEditor(cheat.Description),
			map[bool]int{true: editFlags.weight, false: cheat.Weight}[cmd.Flags().Changed("weight")],
		)

		utils.Render(
			"Edited cheat {BOLD}{GREEN}{name}{RESET}",
			map[string]string{"name": cheat.Name},
		)
	},
}
