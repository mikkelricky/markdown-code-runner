package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/mikkelricky/go-markdown-code-runner/codeblock"
	"github.com/spf13/cobra"
)

const defaultFilename = "README.md"

func showCollection() {
	fmt.Printf("%d block(s) found\n", len(collection.Blocks()))

	for index := range collection.Blocks() {
		if index > 0 {
			fmt.Println(strings.Repeat("-", 120))
			fmt.Println()
		}

		showBlock(strconv.Itoa(index), index)
	}
}

var headerTransformer = text.Transformer(func(val interface{}) string {
	return text.Bold.Sprint(val)
})

func showBlock(id string, index int) {
	block, err := collection.Get(id)
	if err != nil {
		log.Fatal(err)
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
		fmt.Println("Execute this block:")
		fmt.Println()

		if name == "" {
			name = id
		}
		cmd := []string{
			mainScript,
			fmt.Sprintf("--execute %s", name),
		}
		fmt.Printf("%s\n", strings.Join(cmd, " "))
	}
	fmt.Println()
}

func executeBlock() {
	block, err := collection.Get(execute)
	if err != nil {
		log.Fatal(err)
	}
	options := map[string]string{
		"echo":    echo,
		"verbose": strconv.FormatBool(verbose),
	}
	if _, err := block.Execute(options); err != nil {
		log.Fatal(err)
	}
}

var (
	verbose    bool
	execute    string
	show       string
	echo       string
	collection codeblock.CodeBlockCollection
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

		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			fi, _ := os.Stdin.Stat()
			if (fi.Mode() & os.ModeCharDevice) == 0 {
				collection, err = codeblock.ParseReader(bufio.NewReader(os.Stdin))
				if err != nil {
					log.Fatal(err)
				}
			} else {
				filename := defaultFilename
				if len(args) == 1 {
					filename = args[0]
				}
				collection, err = codeblock.ParsePath(filename)
				if err != nil {
					log.Fatal(err)
				}
			}

			if execute != "" {
				executeBlock()
			} else if show != "" {
				showBlock(show, -1)
			} else {
				showCollection()
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

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&show, "show", "s", "", "name or index of a code block to show")
	rootCmd.PersistentFlags().StringVarP(&execute, "execute", "e", "", "name or index of a code block to execute")
	rootCmd.PersistentFlags().StringVarP(&echo, "echo", "", "", "echo shell statements and prepend with the value of this flag, e.g. --echo='> '")
}
