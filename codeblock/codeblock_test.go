package codeblock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstituteBlock(t *testing.T) {
	testCases := []struct {
		content       string
		substitutions map[string]string

		expected string
	}{
		{
			`echo "Hello «name»!"
`,
			map[string]string{},
			`echo "Hello «name»!"
`,
		},

		{
			`echo "Hello «name»!"
`,
			map[string]string{
				"«name»": "Mikkel",
			},
			`echo "Hello Mikkel!"
`,
		},

		{
			`echo "Hey «name», is your name really «name»?"
`,
			map[string]string{
				"«name»": "Mikkel",
			},
			`echo "Hey Mikkel, is your name really Mikkel?"
`,
		},

		{
			`p := (%x%, %y%)
`,
			map[string]string{
				"%x%": "3",
				"%y%": "14",
			},
			`p := (3, 14)
`,
		},
	}

	for _, testCase := range testCases {
		block := createCodeBlock("``` test", testCase.content, "```")
		actual := block.Substitute(testCase.substitutions)

		assert.Equal(t, testCase.expected, actual)
	}
}
