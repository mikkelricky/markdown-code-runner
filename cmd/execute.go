package cmd

import (
	"fmt"
	"strconv"

	"github.com/mikkelricky/go-markdown-code-runner/codeblock"
	"github.com/spf13/cobra"
)

func executeBlock(collection codeblock.CodeBlockCollection, id string) (string, error) {
	block, err := collection.Get(id)
	if err != nil {
		return "", err
	}
	options := map[string]string{
		"echo":    echo,
		"verbose": strconv.FormatBool(verbose),
	}
	return block.Execute(options)
}

// executeCmd represents the execute command
var (
	echo       string
	executeCmd = &cobra.Command{
		Use:   "execute [name...]",
		Short: "Execude code blocks",
		Long: fmt.Sprintf(`Execute one or more code blocks.

Examples:

%[2]s execute test
%[2]s execute test --verbose
%[2]s execute test --verbose --file codeblock/testdata/tests.md
%[2]s execute test --verbose --echo '👉 '
`, appName, mainScript),
		Run: func(cmd *cobra.Command, args []string) {
			collection, err := readCollection()
			check(err)

			for _, arg := range args {
				_, err = executeBlock(*collection, arg)
				check(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(executeCmd)

	executeCmd.Flags().StringVarP(&echo, "echo", "", "", "echo shell statements and prepend with the value of this flag, e.g. --echo='> '")
}
