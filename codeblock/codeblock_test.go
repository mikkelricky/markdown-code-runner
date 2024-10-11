package codeblock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstituteBlock(t *testing.T) {
	testCases := []struct {
		infoString    string
		content       string
		substitutions map[string]string

		expected string
	}{
		{
			"shell",
			`echo "Hello «name»!"
`,
			map[string]string{},
			`echo "Hello «name»!"
`,
		},

		{
			"shell substitutions='«name»: Mikkel'",
			`echo "Hello «name»!"
`,
			map[string]string{},
			`echo "Hello Mikkel!"
`,
		},

		{
			"shell substitutions='«name»: Mikkel'",
			`echo "Hello «name»!"
`,
			map[string]string{
				"«name»": "World",
			},
			`echo "Hello World!"
`,
		},

		{
			"shell",
			`echo "Hello «name»!"
`,
			map[string]string{
				"«name»": "Mikkel",
			},
			`echo "Hello Mikkel!"
`,
		},

		{
			"shell",
			`echo "Hey «name», is your name really «name»?"
`,
			map[string]string{
				"«name»": "Mikkel",
			},
			`echo "Hey Mikkel, is your name really Mikkel?"
`,
		},

		{
			"shell",
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
		block := createCodeBlock("``` "+testCase.infoString, testCase.content, "```")
		actual, err := block.Substitute(testCase.substitutions)
		assert.Nil(t, err)
		assert.Equal(t, testCase.expected, actual)
	}
}
