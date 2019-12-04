package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"cheat/cli/db"
	"cheat/cli/utils"
)

var getDescription string = strings.TrimSpace(`
Get a cheat info by id
`)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Get cheat info",
	Long:    getDescription,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cheat := db.GetCheatByID(args[0])

		// header
		utils.Render(
			"    {BOLD}{BLUE}{command}{RESET}\n",
			map[string]string{"command": cheat.Command},
		)

		// body
		utils.Render(
			cheat.Description,
			nil,
		)

		// footer
		utils.Render(
			"\n{GREY}{BOLD}{id}{RESET}: {GREY}{name}{RESET}",
			map[string]string{
				"id":   cheat.ID,
				"name": cheat.Name,
			},
		)
	},
}
