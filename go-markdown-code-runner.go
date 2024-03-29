package main

import (
	"bufio"
	"fmt"
	"go-markdown-code-runner/codeblock"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "«filename»")
		os.Exit(1)
	}

	filename := args[0]

	file, err := os.Open(filename)
	check(err)

	defer file.Close()

	reader := bufio.NewReader(file)
	blocks := codeblock.Parse(reader)

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
