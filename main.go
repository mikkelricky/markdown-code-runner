package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/rimi-itk/go-markdown-code-runner/codeblock"
)

func main() {
	var id string
	var execute bool
	var echo string
	var verbose bool
	flag.StringVar(&id, "id", "", "Name or index of code block to run")
	flag.BoolVar(&execute, "run", false, "Execute block")
	flag.StringVar(&echo, "echo", "", "Echo shell statements and prepend with the value of this flag, e.g. --echo='> '")
	flag.BoolVar(&verbose, "verbose", false, "Makes command verbose")
	flag.Parse()

	filename := "README.md"
	args := flag.Args()
	if len(args) > 0 {
		filename = args[0]
	}

	collection := codeblock.ParsePath(filename)

	if id != "" {
		block, err := collection.Get(id)
		if err != nil {
			log.Fatal(err)
		}

		if execute {
			options := map[string]string{
				"echo":    echo,
				"verbose": strconv.FormatBool(verbose),
			}
			if err := block.Execute(options); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Print(block)
		}
	} else {
		fmt.Printf("%d block(s) found\n", collection.Len())

		headerTransformer := text.Transformer(func(val interface{}) string {
			return text.Bold.Sprint(val)
		})

		for index, block := range collection.Blocks() {
			name := block.GetName()

			header := fmt.Sprintf("#%d", index)
			if name != "" {
				header += fmt.Sprintf(": %s", name)
			}

			if index > 0 {
				fmt.Println()
				fmt.Println(strings.Repeat("-", 80))
			}

			fmt.Println()
			fmt.Println(headerTransformer(header))

			fmt.Println()
			fmt.Print(block)

			if verbose {
				fmt.Println()
				fmt.Println("Run this block:")
				fmt.Println()

				id = fmt.Sprintf("%d", index)
				if name != "" {
					id = name
				}
				cmd := []string{
					flag.CommandLine.Name(),
					fmt.Sprintf("--id=%s", id),
					"--run",
				}
				fmt.Printf("%s\n", strings.Join(cmd, " "))
			}
		}
	}
}
