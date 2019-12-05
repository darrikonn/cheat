package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"cheat/cli/db"
	"cheat/cli/utils"
)

var (
	getFlags = &struct {
		ignoreCase bool
	}{}
)

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&getFlags.ignoreCase, "ignore-case", "i", false, "Case insensitive search")
	getCmd.SetUsageTemplate(createUsageTemplate("cheat [regex] get [flags]"))
}

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Get cheat info",
	Long: strings.TrimSpace(`
Get a cheat info by name
`),

	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cheat := db.GetCheatByName(args[0], getFlags.ignoreCase)

		// Header
		utils.Render(
			"    {BOLD}{BLUE}{name}{RESET}\n",
			map[string]string{"name": cheat.Name},
		)

		// Body
		utils.Render(
			strings.TrimSpace(cheat.Description),
			nil,
		)

		// Footer
		utils.Render(
			"\n{BOLD}{GREY}{name}{RESET}{GREY}{split} {description}{RESET}",
			map[string]string{
				"name":        cheat.Name,
				"split":       map[bool]string{true: " :", false: ""}[cheat.Description != ""],
				"description": utils.GetFirstLine(cheat.Description),
			},
		)
	},
}
