package codeblock

import (
	"fmt"
	"strconv"
)

type CodeBlockCollection struct {
	blocks      []CodeBlock
	namedBlocks map[string]CodeBlock
}

func NewCodeBlockCollection(blocks []CodeBlock) CodeBlockCollection {
	namedBlocks := map[string]CodeBlock{}
	for _, block := range blocks {
		if name := block.GetName(); name != "" {
			namedBlocks[name] = block
		}
	}

	return CodeBlockCollection{
		blocks,
		namedBlocks,
	}
}

func (collection CodeBlockCollection) Len() int {
	return len(collection.blocks)
}

func (collection CodeBlockCollection) Blocks() []CodeBlock {
	return collection.blocks
}

// Get returns a code block identified by it's name or index in the collection
func (collection CodeBlockCollection) Get(id string) (*CodeBlock, error) {
	if block, found := collection.namedBlocks[id]; found {
		return &block, nil
	}

	index, err := strconv.Atoi(id)
	if err != nil {
		index = -1
	}
	if 0 <= index && index < collection.Len() {
		return &collection.blocks[index], nil
	}

	return nil, fmt.Errorf("cannot find block %s", id)
}
