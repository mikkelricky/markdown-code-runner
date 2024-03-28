package codeblock

import (
	"bufio"
	"io"
	"regexp"
)

func Parse(reader *bufio.Reader) []CodeBlock {
	var blocks []CodeBlock

	// https://github.github.com/gfm/#fenced-code-blocks
	codeBlockStart, _ := regexp.Compile("^ {0,3}```(?P<infoString>[^`]+)")
	codeBlockEnd, _ := regexp.Compile("^ {0,3}```")

	var block CodeBlock
	var inCodeBlock bool = false

	for {
		l, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		line := string(l)

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

	return blocks
}
