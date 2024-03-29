package main

import (
	"bufio"
	"flag"
	"fmt"
	"go-markdown-code-runner/codeblock"
	"log"
	"os"
	"strconv"
	"strings"
)

func findBlock(blocks []codeblock.CodeBlock, id string) (*codeblock.CodeBlock, error) {
	blockIndex, err := strconv.Atoi(id)
	if err != nil {
		blockIndex = -1
	}
	for index, block := range blocks {
		if index == blockIndex || id == block.GetName() {
			return &block, nil
		}
	}
	return nil, fmt.Errorf("cannot find block with id %s", id)
}

func main() {
	var id string
	var execute bool
	var echo string
	flag.StringVar(&id, "id", "", "Name or index of code block to run")
	flag.BoolVar(&execute, "run", false, "Execute block")
	flag.StringVar(&echo, "echo", "", "Echo shell statements and prepend with the value of this flag, e.g. --echo='> '")
	flag.Parse()

	filename := "README.md"
	args := flag.Args()
	if len(args) > 0 {
		filename = args[0]
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	blocks := codeblock.Parse(reader)

	if len(id) > 0 {
		block, err := findBlock(blocks, id)
		if err != nil {
			log.Fatal(err)
		}

		if execute {
			options := map[string]string{
				"echo": echo,
			}
			if err := block.Execute(options); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Print(block)
		}
	} else {
		fmt.Printf("%d block(s) found\n", len(blocks))

		for index, block := range blocks {
			fmt.Println()
			fmt.Println(strings.Repeat("=", 80))
			fmt.Printf("%d", index)
			name := block.GetName()
			if len(name) > 0 {
				fmt.Printf(": %s\n", name)
			}
			fmt.Println()
			fmt.Println(strings.Repeat("-", 80))
			fmt.Print(block)
			fmt.Println(strings.Repeat("=", 80))
		}
	}
}
