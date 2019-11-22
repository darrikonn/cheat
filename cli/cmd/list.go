package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"cheat/cli/db"
	"cheat/cli/utils"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all your cheats",
	Long:    `List all your cheats`,
	Run: func(cmd *cobra.Command, args []string) {
		cheats := db.GetCheats()

		// header
		utils.Render(
			"    ordered by weight\n",
			nil,
		)

		// body
		for _, cheat := range cheats {
			utils.Render(
				"{BOLD}{id}{RESET}: {BOLD}{BLUE}{command}{RESET} {GREY}➞{RESET} {name}",
				map[string]string{
					"id":      cheat.ID,
					"command": cheat.Command,
					"name":    cheat.Name,
				},
			)
		}

		// footer
		utils.Render(
			"\n{GREY}{BOLD}{items}{RESET}{GREY} {plural} stored",
			map[string]string{
				"items":  strconv.Itoa(len(cheats)),
				"plural": utils.SingularOrPlural("cheat", len(cheats)),
			},
		)
	},
}
