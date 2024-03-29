package codeblock

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertInfoStringsEqual(t assert.TestingT, this InfoString, that InfoString) bool {

	if !assert.Equal(t, this.name, that.name, "info strings should have the same name") {
		return false
	}

	if !assert.Equal(t, len(this.properties), len(that.properties), "info strings should have the same number of properties") {
		return false
	}

	for name, value := range this.properties {
		thatValue, found := that.properties[name]
		if !assert.True(t, found, "property %q should exists in both info strings", name) {
			fmt.Printf("\n\n%s\n%s\n\n", this, that)
			return false
		}
		if !assert.Equalf(t, value, thatValue, "property %q should have value %q; found %q found", name, value, thatValue) {
			return false
		}
	}

	return true
}

func TestParseInfoString(t *testing.T) {
	testCases := []struct {
		input     string
		expected  InfoString
		errorText string
	}{
		{
			"shell name=my-test debug",
			InfoString{
				name: "shell",
				properties: map[string]string{
					"name":  "my-test",
					"debug": "",
				},
			},
			"",
		},

		{
			" shell name='my-test' debug=no",
			InfoString{
				name: "shell",
				properties: map[string]string{
					"name":  "my-test",
					"debug": "no",
				},
			},
			"",
		},

		{
			"   ",
			InfoString{},
			"",
		},

		{
			"//",
			InfoString{},
			"error parsing info string `//`",
		},

		{
			"  a-b-c  ",
			InfoString{
				name: "a-b-c",
			},
			"",
		},

		{
			" shell  ",
			InfoString{
				name: "shell",
			},
			"",
		},
	}

	for _, testCase := range testCases {
		actual, err := ParseInfoString(testCase.input)

		if testCase.errorText != "" {
			assert.NotNil(t, err)
			assert.EqualError(t, err, testCase.errorText)
		} else {
			assert.Nil(t, err)
		}

		assertInfoStringsEqual(t, testCase.expected, actual)
	}
}
