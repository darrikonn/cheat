package cmd

import (
	"cheat/cli/db"
	"cheat/cli/utils"
	"strings"

	"github.com/spf13/cobra"
)

var (
	addFlags = &struct {
		weight int
	}{}
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&addFlags.weight, "weight", "w", 0, "weight of the cheat; used for sorting query results")
	addCmd.SetUsageTemplate(createUsageTemplate("cheat [regex] add [flags]"))
}

var addCmd = &cobra.Command{
	Use:        "add",
	Aliases:    []string{"a"},
	SuggestFor: []string{"ad", "addd", "aad", "aa"},
	Short:      "Add a new cheat",
	Long: strings.TrimSpace(`
Add a new cheat to your cheatsheet. You'll be prompted for
the cheat's "description" in your preferred editor.
`),

	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := db.AddCheat(args[0], utils.GetUserInputFromEditor("<summary>\n\n<description>"), addFlags.weight)
		utils.Render(
			"Created cheat {BOLD}{GREEN}{name}{RESET}",
			map[string]string{"name": name},
		)
	},
}
