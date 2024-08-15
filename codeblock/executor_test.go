package codeblock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExcecuteBlock(t *testing.T) {
	block := createCodeBlock(
		"``` sh",
		`echo -n " sh"
`,
		"```",
	)
	expected := ` sh`
	actual, err := block.Run(map[string]string{}, map[string]string{})
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
