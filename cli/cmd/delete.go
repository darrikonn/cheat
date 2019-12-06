package cmd

import (
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"cheat/cli/db"
	"cheat/cli/exceptions"
	"cheat/cli/utils"
)

var (
	deleteFlags = &struct {
		yes        bool
		ignoreCase bool
	}{}
)

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&deleteFlags.yes, "yes", "y", false, "Skip prompt")
	deleteCmd.Flags().BoolVarP(&deleteFlags.ignoreCase, "ignore-case", "i", false, "Case insensitive search")
	deleteCmd.SetUsageTemplate(createUsageTemplate("cheat [regex] delete [flags]"))
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete a cheat",
	Long: strings.TrimSpace(`
Delete a cheat from your cheatsheet.
`),

	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cheat := db.GetCheatByName(args[0], deleteFlags.ignoreCase)
		if !deleteFlags.yes {
			utils.Render(
				"    {BOLD}{name}{RESET}: {description}\n",
				map[string]string{
					"name":        cheat.Name,
					"description": utils.GetFirstLine(cheat.Description),
				},
			)

			prompt := promptui.Prompt{
				Label:     "Are you sure you want to delete <Cheat: " + cheat.Name + ">",
				IsConfirm: true,
			}
			_, err := prompt.Run()
			if err != nil {
				panic(exceptions.Abort(""))
			}
		}

		db.DeleteCheat(cheat.Name)
		utils.Render(
			"{RED}Deleted{RESET} {BOLD}{name}{RESET}",
			map[string]string{
				"name": cheat.Name,
			},
		)
	},
}
