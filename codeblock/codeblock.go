package codeblock

import (
	"fmt"
	"strings"
)

type CodeBlock struct {
	infoString string
	content    []string
}

func New(infoString string) CodeBlock {
	return CodeBlock{infoString: infoString, content: make([]string, 0)}
}

func (block CodeBlock) String() string {
	return fmt.Sprintf("```%v\n%v```\n", block.infoString, block.GetContent())
}

func (block CodeBlock) GetLanguage() string {
	return block.infoString
}

func (block CodeBlock) GetContent() string {
	return strings.Join(block.content[:], "\n")
}

func (block *CodeBlock) AddLine(line string) {
	block.content = append(block.content, line)
}
