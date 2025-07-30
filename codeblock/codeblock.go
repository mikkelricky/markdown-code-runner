package codeblock

import (
	"fmt"
	"log"
	"maps"
	"strings"

	"al.essio.dev/pkg/shellescape"

	"github.com/goccy/go-yaml"
)

type CodeBlock struct {
	infoString     InfoString
	content        []string
	codeBlockStart string
	codeBlockEnd   string
}

func NewCodeBlock(codeBlockStart string, content []string, codeBlockEnd string) CodeBlock {
	match := codeBlockStartPattern.FindStringSubmatch(codeBlockStart)
	if len(match) < 2 {
		log.Fatalf("invalid code block start: %s", codeBlockStart)
	}
	infoString, err := ParseInfoString(match[2])
	if err != nil {
		log.Fatal(err)
	}

	return CodeBlock{
		infoString,
		content,
		codeBlockStart,
		codeBlockEnd,
	}
}

func (block CodeBlock) String() string {
	return fmt.Sprintf("%s\n%s%s\n", block.codeBlockStart, block.GetContent(), block.codeBlockEnd)
}

func (block CodeBlock) GetLanguage() string {
	return block.infoString.GetName()
}

func (block CodeBlock) GetName() string {
	return block.infoString.properties["name"]
}

func (block CodeBlock) GetContent() string {
	return strings.Join(block.content[:], "\n")
}

func (block *CodeBlock) AddLine(line string) {
	block.content = append(block.content, line)
}

func (block *CodeBlock) GetSubstitutions(substitutions map[string]string) (map[string]string, error) {
	value := block.infoString.GetProperty("substitutions")

	var blockSubstitutions map[string]string
	err := yaml.Unmarshal([]byte(value), &blockSubstitutions)
	if err != nil {
		return map[string]string{}, fmt.Errorf("error parsing substitutions %v: %s", shellescape.Quote(value), err.Error())
	} else if len(blockSubstitutions) == 0 {
		blockSubstitutions = map[string]string{}
	}

	maps.Copy(blockSubstitutions, substitutions)

	return blockSubstitutions, nil
}

func (block *CodeBlock) Substitute(substitutions map[string]string) (string, error) {
	content := block.GetContent()
	blockSubstitutions, err := block.GetSubstitutions(substitutions)
	if err != nil {
		return "", err
	}
	for from, to := range blockSubstitutions {
		content = strings.ReplaceAll(content, from, to)
	}

	return content, nil
}
