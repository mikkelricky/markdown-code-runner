package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/mikkelricky/markdown-code-runner/codeblock"
	"github.com/spf13/cobra"
)

func runBlock(collection codeblock.CodeBlockCollection, id string) error {
	block, err := collection.Get(id)
	if err != nil {
		return err
	}
	options := map[string]string{
		"echo":    echo,
		"verbose": strconv.FormatBool(verbose),
	}

	return block.Run(options, substitutions, "")
}

// runCmd represents the run command
var (
	echo   string
	runCmd = &cobra.Command{
		Use:   "run name...",
		Short: "Run code blocks",
		Long:  "Run one or more code blocks.",
		Example: fmt.Sprintf(`
%[2]s run test
%[2]s run test --verbose
%[2]s run test --verbose --file codeblock/testdata/tests.md
%[2]s run test --verbose --echo '👉 '`, appName, mainScript),
		Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
		Run: func(cmd *cobra.Command, args []string) {
			ParseSubstitutions()
			collection, err := readCollection()
			check(err)

			for _, arg := range args {
				err = runBlock(*collection, arg)
				if err != nil {
					if exitErr, ok := err.(*exec.ExitError); ok {
						os.Exit(exitErr.ExitCode())
					} else {
						check(err)
					}
				}
			}
		},
		// https://github.com/spf13/cobra/blob/main/site/content/completions/_index.md#dynamic-completion-of-nouns
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			// if len(args) != 0 {
			// 	return nil, cobra.ShellCompDirectiveNoFileComp
			// }

			collection, err := readCollection()
			if err != nil {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			blockNames := []string{}

			// @todo Optimize!
		OUTER:
			for _, block := range collection.Blocks() {
				name := block.GetName()
				for _, arg := range args {
					if arg == name {
						continue OUTER
					}
				}
				blockNames = append(blockNames, name)
			}

			return blockNames, cobra.ShellCompDirectiveNoFileComp
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&echo, "echo", "", "", "echo shell statements and prepend with the value of this flag, e.g. --echo='> '")
}
