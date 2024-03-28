package codeblock

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func Parse(reader *bufio.Reader) CodeBlockCollection {
	var blocks []CodeBlock

	// https://github.github.com/gfm/#fenced-code-blocks
	codeBlockStart, _ := regexp.Compile("^ {0,3}```(?P<infoString>[^`]+)")
	codeBlockEnd, _ := regexp.Compile("^ {0,3}```")

	var block CodeBlock
	var inCodeBlock bool = false

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		match := codeBlockStart.FindStringSubmatch(line)
		if len(match) > 0 {
			block = NewCodeBlock(match[1])
			inCodeBlock = true
		} else if codeBlockEnd.MatchString(line) {
			block.AddLine("")
			blocks = append(blocks, block)
			inCodeBlock = false
		} else if inCodeBlock {
			block.AddLine(line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

	return NewCodeBlockCollection(blocks)
}

func ParseFile(file *os.File) CodeBlockCollection {
	return Parse(bufio.NewReader(file))
}

func ParsePath(path string) CodeBlockCollection {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return ParseFile(file)
}

func ParseString(text string) CodeBlockCollection {
	return Parse(bufio.NewReader(strings.NewReader(text)))
}
