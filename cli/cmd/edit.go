package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"cheat/cli/cmd/exceptions"
	"cheat/cli/db"
	"cheat/cli/models"
	"cheat/cli/utils"
)

var editDescription string = strings.TrimSpace(`
Edit a cheat's "command", "name", and/or "description".
Pass the "id" of the cheat to be edited.
`)

var (
	editFlags = &struct {
		name        string
		description string
		weight      int
	}{}
)

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().StringVarP(&editFlags.name, "name", "n", "", "name of the cheat")
	editCmd.Flags().StringVarP(&editFlags.description, "description", "d", "", "description of the cheat")
	editCmd.Flags().IntVarP(&editFlags.weight, "weight", "w", 0, "weight of the cheat; used for sorting query results")
}

func validateViewModel(cheat models.Cheat, flags *pflag.FlagSet) {
	nameChanged, descriptionChanged, weightChanged := true, true, true
	if !flags.Changed("name") {
		editFlags.name = cheat.Name
		nameChanged = false
	}
	if !flags.Changed("description") {
		editFlags.description = cheat.Description
		descriptionChanged = false
	}
	if !flags.Changed("weight") {
		editFlags.weight = cheat.Weight
		weightChanged = false
	}

	if !nameChanged && !descriptionChanged && !weightChanged {
		panic(exceptions.Abort("Must provide at least one of \"name\", \"description\", \"weight\" when editing a cheat"))
	}
}

var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e"},
	Short:   "Edit a cheat",
	Long:    editDescription,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cheat := db.GetCheatByID(args[0])

		validateViewModel(cheat, cmd.Flags())
		db.EditCheat(cheat.ID, editFlags.name, editFlags.description, editFlags.weight)

		utils.Render(
			"Edited cheat {BOLD}{GREEN}{id}{RESET}",
			map[string]string{"id": cheat.ID},
		)
	},
}
