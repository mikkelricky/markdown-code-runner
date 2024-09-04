package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alessio/shellescape"
	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

const appName = "markdown-code-runner"
const defaultFilename = "README.md"
const version = "v1.2.1"
const ARG_FILENAME = "file"
const FILENAME_STDIN = "-"
const ARG_SUBSTITUTIONS = "substitutions"
const ARG_VERBOSE = "verbose"

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	verbose          bool
	filename         string
	argSubstitutions string
	substitutions    = map[string]string{}

	mainScript = func() string {
		// 	If the app is run with `go run`, the executable will (probably) be in the temporary folder.
		if executable, err := os.Executable(); err == nil && !strings.HasPrefix(executable, os.TempDir()) {
			// We think that we're not being run with `go run`.
			return os.Args[0]
		}

		return "go run github.com/mikkelricky/markdown-code-runner@latest"
	}()

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:     os.Args[0],
		Short:   "Show and run code blocks in Markdown files",
		Args:    cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
		Version: version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := yaml.Unmarshal([]byte(argSubstitutions), &substitutions)
			if err != nil {
				if verbose {
					log.Fatalf("Error parsing substitutions %v; a valid YAML object expected.\n\nError: %v", shellescape.Quote(argSubstitutions), err)
				} else {
					log.Fatalf("Error parsing substitutions %v; a valid YAML object expected (use --%v for more detailed error message).", shellescape.Quote(argSubstitutions), ARG_VERBOSE)
				}
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func ParseSubstitutions() {
	// err := yaml.Unmarshal([]byte(argSubstitutions), &substitutions)
	// if err != nil {
	// 	log.Fatalf("Error parsing substitutions: %v", argSubstitutions)
	// }
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, ARG_VERBOSE, "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&filename, ARG_FILENAME, "f", defaultFilename, fmt.Sprintf("The file to read code blocks from. Use %s to read from stdin.", FILENAME_STDIN))
	// https://github.com/spf13/cobra/blob/main/site/content/completions/_index.md#specify-valid-filename-extensions-for-flags-that-take-a-filename
	rootCmd.RegisterFlagCompletionFunc(ARG_FILENAME, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"md"}, cobra.ShellCompDirectiveFilterFileExt
	})

	rootCmd.PersistentFlags().StringVarP(&argSubstitutions, ARG_SUBSTITUTIONS, "s", "", "Substitutions to apply before show and run. Must be a valid YAML object, e.g. 'id: 87' or '{id: 87, name: \"Mikkel\"}'")
}
