package cmd

import (
  "strconv"

  "github.com/spf13/cobra"

  "cheat/cli/db"
  "cheat/cli/utils"
)

func init() {
  rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
  Use: "search",
  Aliases: []string{"s"},
  Short: "Search your cheats",
  Long: `Search your cheats from a regex`,
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    cheats := db.SearchCheats(args[0])

    // header
    utils.Render(
      "    {search}\n",
      map[string]string{
        "search": args[0],
      },
    )

    // body
    for _, cheat := range cheats {
      utils.Render(
        "{BOLD}{id}{RESET}: {BOLD}{BLUE}{command}{RESET} {GREY}➞{RESET} {name}",
        map[string]string{
          "id": cheat.ID,
          "command": cheat.Command,
          "name": cheat.Name,
        },
      )
    }

    // footer
    utils.Render(
			"\n{GREY}{BOLD}{items}{RESET}{GREY} {plural} match your search",
      map[string]string{
        "items": strconv.Itoa(len(cheats)),
        "plural": utils.SingularOrPlural("cheat", len(cheats)),
      },
    )
  },
}
