package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	db "cheat/cli/db"
	"cheat/cli/exceptions"
	"cheat/cli/utils"
)

var (
	// Used for flags.
	verbose bool

	defaultCommand string = "search"

	rootCmd = &cobra.Command{
		Use:          "cheat",
		Version:      "0.1.0",
		SilenceUsage: true,
		Short:        "Cheat is a personal cheatsheet manager",
		Long: strings.TrimSpace(`
A fast and flexible cheatsheet manager built with
Go. Complete documentation is available at
https://github.com/darrikonn/cheat/api.md
`),
	}
)

// Name returns the command's name: the first word in the use line.
func init() {
	initConfig()
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose output")
	rootCmd.SetUsageTemplate(createUsageTemplate("cheat [regex] [command]"))
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
}

func setArgs() {
	type _Argument struct {
		value string
		index int
	}
	regex := &_Argument{value: "", index: -1}
	target := &_Argument{value: defaultCommand, index: -1}

	// Traverse args to find regex and target
	args := os.Args[1:]
	for i, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			if regex.index < 0 {
				regex = &_Argument{value: arg, index: i}
			} else if target.index < 0 {
				target = &_Argument{value: arg, index: i}
				break
			}
		}
	}

	// Flip regex and target
	regexFound := regex.index >= 0
	targetFound := target.index >= 0
	if regexFound {
		args = utils.RemoveAtIndex(args, regex.index)
		if targetFound {
			args = utils.RemoveAtIndex(args, target.index-1)
		}
	}

	if !regexFound && !targetFound && utils.ContainsAny(args, "--help", "-h") {
		// No command initiated, but help requested
		return
	}

	// Finally set the args
	rootCmd.SetArgs(append([]string{target.value, regex.value}, args...))
}

func errorHandling() {
	err := recover()
	if err != nil {
		switch err.(type) {
		case *exceptions.AbortType:
			utils.Render("Abort!", nil)
			message := err.(*exceptions.AbortType).Error()
			if message != "" {
				utils.Render("{RED}{message}{RESET}", map[string]string{"message": message})
			}
		case *exceptions.CheatExceptionType:
			utils.Render("{RED}Error{RESET}:", nil)
			utils.Render(err.(*exceptions.CheatExceptionType).Error(), nil)
			if verbose && err.(*exceptions.CheatExceptionType).Original() != nil {
				panic(err.(*exceptions.CheatExceptionType).Original())
			}
		default:
			if verbose {
				panic(err)
			} else {
				utils.Render("{RED}Error{RESET}:", nil)
				utils.Render(err.(error).Error(), nil)
			}
		}
		os.Exit(1)
	}
}

func initConfig() {
	// Find home directory.
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		panic(exceptions.CheatException("Could not find home directory", err))
	}

	// Search config in home directory with name ".cheet" (without extension).
	viper.AddConfigPath(homeDirectory)
	viper.SetConfigName(".cheet")

	// Fallback to "vi" for the editor
	viper.SetDefault("editor", utils.GetEnv("EDITOR", "vi"))
	viper.SetDefault("database", "~/.cheetsheet.db")

	// Load config
	_ = viper.ReadInConfig()
}

func createUsageTemplate(usage string) string {
	return `Usage:
  ` + usage + `{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{if eq .Name "` + defaultCommand + `"}}{{rpad "*" .NamePadding }} {{.Short}} (default)` +
		`{{else}}{{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
}

// Execute : executes the root command
// that combines all cli subcommands
func Execute() {
	// Setup the database and defer clean it up
	db.Setup()
	defer db.Cleanup()

	// Set args for our API
	setArgs()

	// Let deferred error handler take care of errors
	defer errorHandling()

	_ = rootCmd.Execute()
}
