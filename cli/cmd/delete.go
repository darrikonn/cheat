package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"cheat/cli/cmd/exceptions"
	"cheat/cli/db"
	"cheat/cli/utils"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a cheat",
	Long: `Delete a cheat from your cheatsheet. Pass the "id" of
         the cheat to be deleted.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cheat := db.GetCheatByID(args[0])
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
