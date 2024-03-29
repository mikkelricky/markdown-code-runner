package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const appName = "go-markdown-code-runner"
const defaultFilename = "README.md"

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	verbose  bool
	filename string

	mainScript = func() string {
		// 	If the app is run with `go run`, the executable will (probably) be in the temporary folder.
		if executable, err := os.Executable(); err == nil && !strings.HasPrefix(executable, os.TempDir()) {
			// We think that we're not being run with `go run`.
			return os.Args[0]
		}

		return "go run github.com/mikkelricky/go-markdown-code-runner@latest"
	}()

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   os.Args[0],
		Short: "Show and run code blocks in Markdown files",
		Args:  cobra.MatchAll(cobra.MaximumNArgs(1), cobra.OnlyValidArgs),
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

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&filename, "file", "f", defaultFilename, "The file to read code blocks from")
}
