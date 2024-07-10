package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/mikkelricky/markdown-code-runner/codeblock"
	"github.com/spf13/cobra"
)

func readCollection() (*codeblock.CodeBlockCollection, error) {
	var (
		collection codeblock.CodeBlockCollection
		err        error
	)
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		collection, err = codeblock.ParseReader(bufio.NewReader(os.Stdin))
		if err != nil {
			return nil, err
		}
	} else {
		collection, err = codeblock.ParsePath(filename)
		if err != nil {
			return nil, err
		}
	}

	return &collection, nil
}

func showCollection(collection codeblock.CodeBlockCollection) {
	fmt.Printf("%d block(s) found\n", len(collection.Blocks()))

	for index := range collection.Blocks() {
		if index > 0 {
			fmt.Println(strings.Repeat("-", 120))
			fmt.Println()
		}

		showBlock(collection, strconv.Itoa(index), index)
	}
}

var headerTransformer = text.Transformer(func(val interface{}) string {
	return text.Bold.Sprint(val)
})

func showBlock(collection codeblock.CodeBlockCollection, id string, index int) error {
	block, err := collection.Get(id)
	if err != nil {
		return err
	}

	name := block.GetName()
	if index < 0 {
		index, err = strconv.Atoi(id)
		if err != nil {
			index = -1
		}
	}

	headerItems := []string{}
	if name != "" {
		headerItems = append(headerItems, name)
	}
	if index > -1 {
		headerItems = append(headerItems, fmt.Sprintf("(#%d)", index))
	}
	fmt.Println(headerTransformer(strings.Join(headerItems, " ")))
	fmt.Println()

	fmt.Print(block)

	if verbose {
		fmt.Println()
		fmt.Println("Run this block:")
		fmt.Println()

		if name == "" {
			name = id
		}
		cmd := []string{
			mainScript,
			"run",
			name,
		}
		fmt.Printf("%s\n", strings.Join(cmd, " "))
	}
	fmt.Println()

	return nil
}

// showCmd represents the show command
var (
	showCmd = &cobra.Command{
		Use:   "show [name...]",
		Short: "Show code blocks",
		Long: fmt.Sprintf(`Show all or select code blocks.

Examples:

%[2]s show test 89
%[2]s show test --verbose
%[2]s show test --verbose --file codeblock/testdata/tests.md
%[2]s show test --verbose --echo '👉 '
`, appName, mainScript),

		Run: func(cmd *cobra.Command, args []string) {
			collection, err := readCollection()
			check(err)

			if len(args) > 0 {
				for _, arg := range args {
					err = showBlock(*collection, arg, -1)
					check(err)
				}
			} else {
				showCollection(*collection)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(showCmd)
}
