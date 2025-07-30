package codeblock

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFileBootstrap(t *testing.T) {
	text := `~~~shell
~~~
`
	collection, _ := ParseString(text)

	assert.Len(t, collection.Blocks(), 1)
	assert.Equal(t, text, collection.Blocks()[0].String())
}

func TestParseFile(t *testing.T) {
	path := "testdata/tests.md"
	expected := NewCodeBlockCollection([]CodeBlock{
		createCodeBlock(
			"``` sh",
			`echo -n " sh"
`,
			"```",
		),

		createCodeBlock(
			"``` shell",
			`echo " shell"
`,
			"```",
		),

		createCodeBlock(
			"```shell",
			`echo "shell"
`,
			"```",
		),

		createCodeBlock(
			"```shell name=test",
			`echo "shell test"
`,
			"```",
		),

		createCodeBlock(
			"```shell",
			``,
			"```",
		),
	})

	actual, err := ParsePath(path)
	if err != nil {
		t.Fatal(err)
	}
	assertCollectionsAreEqual(t, expected, actual)

	_, err = expected.Get("hest")
	assert.NotNil(t, err)

	_, err = expected.Get("test")
	assert.Nil(t, err)
}

func assertCollectionsAreEqual(t *testing.T, c0 CodeBlockCollection, c1 CodeBlockCollection) {
	if assert.Equal(t, len(c0.Blocks()), len(c1.Blocks()), "collection should have the same size") {
		for i0, b0 := range c0.Blocks() {
			b1 := c1.Blocks()[i0]
			assertBlocksAreEqual(t, b0, b1)
		}
	}
}

func assertBlocksAreEqual(t *testing.T, b0 CodeBlock, b1 CodeBlock) {
	// assert.Equal(t, b0.infoString, b1.infoString, "info strings should be equal")
	assert.Equal(t, b0.GetLanguage(), b1.GetLanguage(), "languages should be equal")
	assert.Equal(t, b0.GetContent(), b1.GetContent(), "content should be equal")
	assert.Equal(t, b0.String(), b1.String(), "rendered blocks should be equal")
}

func createCodeBlock(codeBlockStart string, content string, codeBlockEnd string) CodeBlock {
	block := NewCodeBlock(codeBlockStart, strings.Split(content, "\n"), codeBlockEnd)

	return block
}
