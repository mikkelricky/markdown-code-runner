package codeblock

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

func (block CodeBlock) Run(options map[string]string, substitutions map[string]string) error {
	verbose, _ := strconv.ParseBool(options["verbose"])
	echo := options["echo"]
	language := allLanguages[block.GetLanguage()]

	if verbose {
		fmt.Printf("Running code block\n\n%s\n", block)
	}

	cmd, args, err := getShellCommand(language)
	if err != nil {
		return err
	}

	run := func(args []string) error {
		cmd := exec.Command(cmd, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		return cmd.Run()
	}

	code := block.Substitute(substitutions)

	switch language {
	case PHP:
		if !strings.Contains(code, "<?php") {
			code = "<?php\n" + code
		}
		file, err := os.CreateTemp("", "code-runner-php")
		if err != nil {
			return err
		}
		defer os.Remove(file.Name())

		_, err = file.WriteString(code)
		if err != nil {
			return err
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
		args = append(args, "-c", code)
		return run(args)

	case ZSH:
		args = append(args, "-c", code)
		return run(args)

	default:
		return fmt.Errorf("cannot handle language %s", block.GetLanguage())
	}
}
