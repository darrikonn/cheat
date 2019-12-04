package cmd

import (
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"cheat/cli/cmd/exceptions"
	"cheat/cli/db"
	"cheat/cli/utils"
)

var deleteDescription string = strings.TrimSpace(`
Delete a cheat from your cheatsheet. Pass the "id" of
the cheat to be deleted.
`)

var (
	deleteFlags = &struct {
		yes bool
	}{}
)

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&deleteFlags.yes, "yes", "y", false, "Skip prompt")
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete a cheat",
	Long:    deleteDescription,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cheat := db.GetCheatByID(args[0])
		if !deleteFlags.yes {
			utils.Render(
				"    {BOLD}{id}{RESET}: {BOLD}{BLUE}{command}{RESET} {GREY}➞{RESET} {name}\n",
				map[string]string{
					"id":      cheat.ID,
					"command": cheat.Command,
					"name":    cheat.Name,
				},
			)

			prompt := promptui.Prompt{
				Label:     "Are you sure you want to delete <Cheat: " + cheat.ID + ">",
				IsConfirm: true,
			}
			_, err := prompt.Run()
			if err != nil {
				panic(exceptions.Abort)
			}
		}

		db.DeleteCheat(cheat.ID)
		utils.Render(
			"{RED}Deleted{RESET} {BOLD}{id}{RESET}: {command}",
			map[string]string{
				"id":      cheat.ID,
				"command": cheat.Command,
			},
		)
	},
}
