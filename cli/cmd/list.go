package cmd

import (
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"cheat/cli/db"
	"cheat/cli/utils"
)

var listDescription string = strings.TrimSpace(`
List all your cheats
`)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all your cheats",
	Long:    listDescription,
	Run: func(cmd *cobra.Command, args []string) {
		cheats := db.GetCheats()

		// header
		utils.Render(
			"    showing [id] [command] : [name]\n",
			nil,
		)

		// body
		for _, cheat := range cheats {
			utils.Render(
				"{id} {BOLD}{BLUE}{command}{RESET}{split} {BOLD}{name}{RESET}",
				map[string]string{
					"id":      cheat.ID,
					"command": cheat.Command,
					"split":   map[bool]string{true: " :", false: ""}[cheat.Name != ""],
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
