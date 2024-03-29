package codeblock

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type CodeBlock struct {
	infoString     string
	content        []string
	language       string
	attributes     map[string]string
	codeBlockStart string
	codeBlockEnd   string
}

func NewCodeBlock(codeBlockStart string, content []string, codeBlockEnd string) CodeBlock {
	match := codeBlockStartPattern.FindStringSubmatch(codeBlockStart)
	if len(match) < 2 {
		log.Fatalf("invalid code block start: %s", codeBlockStart)
	}
	infoString := match[2]

	r := regexp.MustCompile("^(?P<language>[^ ]+)(?: +(?P<attributes>.+))?")
	match = r.FindStringSubmatch(strings.TrimSpace(infoString))

	language := match[1]
	attributes := map[string]string{}

	if len(match[2]) > 0 {
		r := regexp.MustCompile("(?P<key>[a-z]+)=(?P<value>[^ ]+)")
		matches := r.FindAllStringSubmatch(match[2], -1)
		for _, match := range matches {
			attributes[match[1]] = match[2]
		}
	}

	return CodeBlock{
		infoString,
		content,
		language,
		attributes,
		codeBlockStart,
		codeBlockEnd,
	}
}

func (block CodeBlock) String() string {
	return fmt.Sprintf("%s\n%s%s\n", block.codeBlockStart, block.GetContent(), block.codeBlockEnd)
}

func (block CodeBlock) GetLanguage() string {
	return block.language
}

func (block CodeBlock) GetAttributes() map[string]string {
	return block.attributes
}

func (block CodeBlock) GetName() string {
	return block.GetAttributes()["name"]
}

func (block CodeBlock) GetContent() string {
	return strings.Join(block.content[:], "\n")
}

func (block *CodeBlock) AddLine(line string) {
	block.content = append(block.content, line)
}
