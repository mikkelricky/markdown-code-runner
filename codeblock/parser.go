package codeblock

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

// https://github.github.com/gfm/#fenced-code-blocks
var codeBlockStartPattern = regexp.MustCompile("^ {0,3}(?P<fence>```|~~~)(?P<infoString>[^`]+)")
var codeBlockEndPattern = regexp.MustCompile("^ {0,3}(?P<fence>```|~~~)")

func ParseReader(reader *bufio.Reader) (CodeBlockCollection, error) {
	var blocks []CodeBlock

	var inCodeBlock bool = false
	var codeBlockStart string
	var code []string
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		match := codeBlockStartPattern.FindStringSubmatch(line)
		if len(match) > 0 {
			codeBlockStart = line
			code = make([]string, 0)
			inCodeBlock = true
		} else if codeBlockEndPattern.MatchString(line) {
			if inCodeBlock {
				if len(code) > 0 {
					code = append(code, "")
				}
				block := NewCodeBlock(codeBlockStart, code, line)
				blocks = append(blocks, block)
			}
			inCodeBlock = false
		} else if inCodeBlock {
			code = append(code, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return NewCodeBlockCollection([]CodeBlock{}), err
	}

	return NewCodeBlockCollection(blocks), nil
}

func ParseFile(file *os.File) (CodeBlockCollection, error) {
	return ParseReader(bufio.NewReader(file))
}

func ParsePath(path string) (CodeBlockCollection, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return ParseFile(file)
}

func ParseString(text string) (CodeBlockCollection, error) {
	return ParseReader(bufio.NewReader(strings.NewReader(text)))
}
