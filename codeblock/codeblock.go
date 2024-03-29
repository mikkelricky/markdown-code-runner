package codeblock

import (
	"bytes"
	"regexp"
	"strings"
)

type CodeBlock struct {
	infoString string
	content    []string
	language   string
	attributes map[string]string
}

func NewCodeBlock(infoString string) CodeBlock {
	r := regexp.MustCompile("^(?P<language>[^ ]+)(?: +(?P<attributes>.+))?")
	match := r.FindStringSubmatch(strings.TrimSpace(infoString))

	language := match[1]
	if language == "sh" {
		language = "shell"
	}

	attributes := map[string]string{}

	if len(match[2]) > 0 {
		r := regexp.MustCompile("(?P<key>[a-z]+)=(?P<value>[^ ]+)")
		matches := r.FindAllStringSubmatch(match[2], -1)
		for _, match := range matches {
			attributes[match[1]] = match[2]
		}
	}

	return CodeBlock{
		infoString: infoString,
		content:    make([]string, 0),
		language:   language,
		attributes: attributes,
	}
}

func (block CodeBlock) String() string {
	var b bytes.Buffer
	b.WriteString("```")
	b.WriteString(block.language)
	for name, value := range block.GetAttributes() {
		b.WriteString(" ")
		b.WriteString(name)
		b.WriteString("=")
		b.WriteString(value)
	}
	b.WriteString("\n")
	b.WriteString(block.GetContent())
	b.WriteString("```\n")

	return b.String()
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
