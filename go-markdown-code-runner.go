package main

import (
	"bufio"
	"fmt"
	"go-markdown-code-runner/codeblock"
	"io"
	"os"
	"regexp"
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

	// https://github.github.com/gfm/#fenced-code-blocks
	codeBlockStart, _ := regexp.Compile("^ {0,3}```(?P<infoString>[^`]+)")
	codeBlockEnd, _ := regexp.Compile("^ {0,3}```")

	var block codeblock.CodeBlock
	var inCodeBlock bool = false

	for {
		l, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		line := string(l)

		match := codeBlockStart.FindStringSubmatch(line)
		if len(match) > 0 {
			block = codeblock.New(match[1])
			inCodeBlock = true
		} else if codeBlockEnd.MatchString(line) {
			block.AddLine("")
			fmt.Print(block)
			inCodeBlock = false
		} else if inCodeBlock {
			block.AddLine(line)
		}
	}

	// fmt.Printf("len=%d cap=%d %v\n", len(code), cap(code), code)

	// content, err := os.ReadFile(filename)
	// check(err)
	// fmt.Print(string(dat))

	// scanner := bufio.NewScanner(os.Stdin)

	// for scanner.Scan() {
	//	line := scanner.Text()
	//	fmt.Println(line)
	// }

	// if err := scanner.Err(); err != nil {
	//	fmt.Fprintln(os.Stderr, "error:", err)
	//	os.Exit(1)
	// }
}
