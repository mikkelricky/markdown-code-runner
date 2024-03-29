package codeblock

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alessio/shellescape"
	"github.com/bitfield/script"
)

const (
	BASH  = "bash"
	SHELL = "shell"
	SH    = "sh"
	ZSH   = "zsh"
	PHP   = "php"
)

// Cf. https://github.com/github-linguist/linguist/blob/master/lib/linguist/languages.yml
var languageAliases = map[string][]string{
	PHP:   {},
	SHELL: {SH, BASH},
	ZSH:   {},
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

func getShellCommand(language string) (string, []string, error) {
	switch language {
	case PHP:
		return "php", []string{}, nil

	case SHELL:
		return "bash", []string{}, nil

	case ZSH:
		return "zsh", []string{}, nil

	default:
		return "", []string{}, fmt.Errorf("cannot get shell command for language %s", language)
	}
}

func (block CodeBlock) Run(options map[string]string) (string, error) {
	verbose, _ := strconv.ParseBool(options["verbose"])
	echo := options["echo"]
	language := allLanguages[block.GetLanguage()]

	if verbose {
		fmt.Printf("Running code block\n\n%s\n", block)
	}

	cmd, args, err := getShellCommand(language)
	if err != nil {
		return "", err
	}

	run := func(args []string) (string, error) {
		cmdLine := strings.Join([]string{cmd, shellescape.QuoteCommand(args)}, " ")

		return script.Exec(cmdLine).WithStderr(os.Stderr).Tee().String()
	}

	switch language {
	case PHP:
		code := block.GetContent()
		if !strings.Contains(code, "<?php") {
			code = "<?php\n" + code
		}
		file, err := os.CreateTemp("", "code-runner-php")
		if err != nil {
			return "", err
		}
		defer os.Remove(file.Name())

		_, err = file.WriteString(code)
		if err != nil {
			return "", err
		}

		args = append(args, "-f", file.Name())
		return run(args)

	case SHELL:
		if len(echo) > 0 {
			args = append(args, "-x")
			// @see `-x` on https://www.gnu.org/software/bash/manual/html_node/The-Set-Builtin.html
			// @see https://github.com/bitfield/script/issues/80
			os.Setenv("PS4", echo)
		}
		args = append(args, "-c", block.GetContent())
		return run(args)

	case ZSH:
		args = append(args, "-c", block.GetContent())
		return run(args)

	default:
		return "", fmt.Errorf("cannot handle language %s", block.GetLanguage())
	}
}
