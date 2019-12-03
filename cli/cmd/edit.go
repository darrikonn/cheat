package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"cheat/cli/cmd/exceptions"
	"cheat/cli/db"
	"cheat/cli/models"
	"cheat/cli/utils"
)

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringVarP(&name, "name", "n", "", "name of the cheat")
	editCmd.Flags().StringVarP(&description, "description", "d", "", "description of the cheat")
	editCmd.Flags().IntVarP(&weight, "weight", "w", 0, "weight of the cheat; used for sorting query results")
}

func validateViewModel(cheat models.Cheat, flags *pflag.FlagSet) (string, string, int) {
	nameChanged, descriptionChanged, weightChanged := true, true, true
	if !flags.Changed("name") {
		name = cheat.Name
		nameChanged = false
	}
	if !flags.Changed("description") {
		description = cheat.Description
		descriptionChanged = false
	}
	if !flags.Changed("weight") {
		weight = cheat.Weight
		weightChanged = false
	}

	if !nameChanged && !descriptionChanged && !weightChanged {
		panic(exceptions.Abort("Must provide at least one of \"name\", \"description\", \"weight\" when editing a cheat"))
	} else {
		return name, description, weight
	}
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a cheat",
	Long: `Edit a cheat's "command", "name", and/or "description".
         Pass the "id" of the cheat to be edited.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cheat := db.GetCheatByID(args[0])

		name, description, weight = validateViewModel(cheat, cmd.Flags())
		db.EditCheat(cheat.ID, name, description, weight)

		utils.Render(
			"Edited cheat {BOLD}{GREEN}{id}{RESET}",
			map[string]string{"id": cheat.ID},
		)
	},
}
