package codeblock

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/alessio/shellescape"
	"github.com/bitfield/script"
)

var languageAliases = map[string][]string{
	"shell": {"sh", "bash"},
}

var allLanguages = func() map[string]string {
	languages := map[string]string{}

	for language, aliases := range languageAliases {
		languages[language] = language
		for _, alias := range aliases {
			languages[alias] = language
		}
	}

	return languages
}()

func getShellCommand(language string) (string, error) {
	switch language {
	case "shell":
		return "sh", nil
	default:
		return "", fmt.Errorf("cannot get shell command for language %s", language)
	}
}

func (block CodeBlock) Execute(options map[string]string) error {
	verbose, _ := strconv.ParseBool(options["verbose"])
	echo := options["echo"]
	language := allLanguages[block.GetLanguage()]

	if verbose {
		fmt.Printf("Executing code block\n\n%s\n", block)
	}

	switch language {
	case "shell":
		cmd, err := getShellCommand(language)
		if err != nil {
			log.Fatal(err)
		}
		args := []string{}
		if len(echo) > 0 {
			args = append(args, "-x")
			// @see `-x` on https://www.gnu.org/software/bash/manual/html_node/The-Set-Builtin.html
			// @see https://github.com/bitfield/script/issues/80
			os.Setenv("PS4", echo)
		}
		args = append(args, "-c", block.GetContent())

		cmdLine := strings.Join([]string{cmd, shellescape.QuoteCommand(args)}, " ")
		script.Exec(cmdLine).WithStderr(os.Stderr).Stdout()

		return nil

	default:
		return fmt.Errorf("cannot handle language %s", block.GetLanguage())
	}
}
