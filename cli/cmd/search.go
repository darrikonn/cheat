package cmd

import (
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"cheat/cli/db"
	"cheat/cli/utils"
)

var (
	searchFlags = &struct {
		ignoreCase bool
	}{}
)

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolVarP(&searchFlags.ignoreCase, "ignore-case", "i", false, "Case insensitive search")
	searchCmd.SetUsageTemplate(createUsageTemplate("cheat [regex] [flags]"))
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search your cheats from regex",
	Long: strings.TrimSpace(`
Search your cheats from a regex
`),
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cheats := db.SearchCheats(args[0], searchFlags.ignoreCase)

		// Header
		utils.Render(
			"    showing {BOLD}[name] : [summary]{RESET}\n", nil,
		)

		// Body
		for _, cheat := range cheats {
			utils.Render(
				"{BOLD}{BLUE}{name}{RESET}{split} {description}{RESET}",
				map[string]string{
					"name":        cheat.Name,
					"split":       map[bool]string{true: " :", false: ""}[cheat.Description != ""],
					"description": utils.GetFirstLine(cheat.Description),
				},
			)
		}

		// Footer
		utils.Render(
			"\n{GREY}{BOLD}{items}{RESET}{GREY} {plural} stored",
			map[string]string{
				"items":  strconv.Itoa(len(cheats)),
				"plural": utils.SingularOrPlural("cheat", len(cheats)),
			},
		)
	},
}
