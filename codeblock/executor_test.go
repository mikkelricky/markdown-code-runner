package codeblock

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dieOn(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

// https://groups.google.com/g/golang-nuts/c/hVUtoeyNL7Y#c124
func catchOutput(t *testing.T, runnable func() error) (string, string, error) {
	stdoutOriginal, stderrOriginal := os.Stdout, os.Stderr
	defer func() {
		os.Stdout = stdoutOriginal
		os.Stderr = stderrOriginal
	}()

	stdoutReader, stdoutFake, err := os.Pipe()
	dieOn(err, t)
	os.Stdout = stdoutFake

	stderrReader, stderrFake, err := os.Pipe()
	dieOn(err, t)
	os.Stderr = stderrFake

	err = runnable()

	// need to close here, otherwise ReadAll never gets "EOF".
	dieOn(stdoutFake.Close(), t)
	stdoutBytes, err := io.ReadAll(stdoutReader)
	dieOn(stdoutReader.Close(), t)

	dieOn(stderrFake.Close(), t)
	stderrBytes, err := io.ReadAll(stderrReader)
	dieOn(stderrReader.Close(), t)

	return string(stdoutBytes), string(stderrBytes), err
}

func TestExcecuteBlock(t *testing.T) {
	testCases := []struct {
		block          CodeBlock
		stdinInput     string
		expectedStdout string
		expectedStderr string
	}{
		{
			createCodeBlock(
				"``` sh",
				`echo -n " test"
`,
				"```",
			),
			``,
			` test`,
			``,
		},

		{
			createCodeBlock(
				"``` shell",
				`echo "This is written on stdout"
(>&2 echo "This is written on stderr")
`,
				"```",
			),
			``,
			`This is written on stdout
`,
			`This is written on stderr
`,
		},

		{
			createCodeBlock(
				"``` shell",
				`read name
echo "Hello $name"
`,
				"```",
			),
			`Mikkel
`,
			`Hello Mikkel
`,
			``,
		},

		{
			createCodeBlock(
				"``` shell",
				`read given_name
read family_name
echo "Hello $family_name, $given_name"
`,
				"```",
			),
			`Donald
Duck
`,
			`Hello Duck, Donald
`,
			``,
		},
	}

	for _, testCase := range testCases {
		actualStdout, actualStderr, err := catchOutput(t, func() error {
			return testCase.block.Run(map[string]string{}, map[string]string{}, testCase.stdinInput)
		})
		assert.Nil(t, err)
		assert.Equal(t, testCase.expectedStdout, actualStdout, "stdout")
		assert.Equal(t, testCase.expectedStderr, actualStderr, "stderr")
	}
}
