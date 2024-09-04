package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alessio/shellescape"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/mikkelricky/markdown-code-runner/codeblock"
	"github.com/spf13/cobra"
)

func readCollection() (*codeblock.CodeBlockCollection, error) {
	var (
		collection codeblock.CodeBlockCollection
		err        error
	)
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	codeblock.Debug("readCollection; fi.Mode(): %q; os.ModeCharDevice: %q; (fi.Mode() & os.ModeCharDevice): %q", fi.Mode(), os.ModeCharDevice, (fi.Mode() & os.ModeCharDevice))
	codeblock.Debug("readCollection; fi.Mode(): %q; os.ModeCharDevice: %q; (fi.Mode() & os.ModeCharDevice): %q", fi.Mode(), os.ModeCharDevice, (fi.Mode()&os.ModeCharDevice) == 0)
	codeblock.Debug("filename: %q", filename)

	if filename == "-" {
		codeblock.Debug("Reading from stdin …")
		collection, err = codeblock.ParseReader(bufio.NewReader(os.Stdin))
		if err != nil {
			return nil, err
		}
	} else {
		codeblock.Debug("Reading from file %s", filename)
		collection, err = codeblock.ParsePath(filename)
		if err != nil {
			return nil, err
		}
	}

	return &collection, nil
}

func showCollection(collection codeblock.CodeBlockCollection, substitutions map[string]string) {
	fmt.Printf("%d block(s) found\n", len(collection.Blocks()))

	for index := range collection.Blocks() {
		if index > 0 {
			fmt.Println(strings.Repeat("-", 120))
			fmt.Println()
		}

		showBlock(collection, strconv.Itoa(index), index, substitutions)
	}
}

var headerTransformer = text.Transformer(func(val interface{}) string {
	return text.Bold.Sprint(val)
})

func showBlock(collection codeblock.CodeBlockCollection, id string, index int, substitutions map[string]string) error {
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

	if block.Substitute(substitutions) != block.GetContent() {
		fmt.Println()
		fmt.Println("With substitutions")
		fmt.Println()
		fmt.Print(block.Substitute(substitutions))
		fmt.Println()
	}

	if verbose {
		fmt.Println()
		fmt.Println("Run this block:")
		fmt.Println()

		if name == "" {
			name = id
		}
		cmd := []string{
			mainScript,
		}

		// Global arguments
		if filename != "" {
			cmd = append(cmd, "--"+ARG_FILENAME, filename)
		}

		cmd = append(cmd, "run")

		if strings.HasPrefix(name, "-") {
			if argSubstitutions != "" {
				cmd = append(cmd, "--"+ARG_SUBSTITUTIONS, shellescape.Quote(argSubstitutions))
			}
			cmd = append(cmd, "--", name)
		} else {
			cmd = append(cmd, name)
			if argSubstitutions != "" {
				cmd = append(cmd, "--"+ARG_SUBSTITUTIONS, shellescape.Quote(argSubstitutions))
			}
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
			ParseSubstitutions()
			collection, err := readCollection()
			check(err)

			if len(args) > 0 {
				for _, arg := range args {
					err = showBlock(*collection, arg, -1, substitutions)
					check(err)
				}
			} else {
				showCollection(*collection, substitutions)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(showCmd)
}
