package codeblock

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var languages = map[string]string{
	"sh": "shell",
}

type CodeBlock struct {
	infoString string
	content    []string
	language   string
	attributes map[string]string
}

func NewCodeBlock(infoString string) CodeBlock {
	r := regexp.MustCompile("^(?P<language>[^ ]+)(?: +(?P<attributes>.+))?")
	match := r.FindStringSubmatch(strings.TrimSpace(infoString))

	language := match[1]
	if language == "sh" {
		language = "shell"
	}

	attributes := map[string]string{}

	if len(match[2]) > 0 {
		r := regexp.MustCompile("(?P<key>[a-z]+)=(?P<value>[^ ]+)")
		matches := r.FindAllStringSubmatch(match[2], -1)
		for _, match := range matches {
			attributes[match[1]] = match[2]
		}
	}

	return CodeBlock{
		infoString: infoString,
		content:    make([]string, 0),
		language:   language,
		attributes: attributes,
	}
}

func (block CodeBlock) String() string {
	var b bytes.Buffer
	b.WriteString("```")
	b.WriteString(block.language)
	for name, value := range block.GetAttributes() {
		b.WriteString(" ")
		b.WriteString(name)
		b.WriteString("=")
		b.WriteString(value)
	}
	b.WriteString("\n")
	b.WriteString(block.GetContent())
	b.WriteString("```\n")

	return b.String()
}

func (block CodeBlock) GetLanguage() string {
	return block.language
}

func (block CodeBlock) GetAttributes() map[string]string {
	return block.attributes
}

func (block CodeBlock) GetName() string {
	return block.GetAttributes()["name"]
}

func (block CodeBlock) GetContent() string {
	return strings.Join(block.content[:], "\n")
}

func (block *CodeBlock) AddLine(line string) {
	block.content = append(block.content, line)
}

func (block CodeBlock) Execute(options map[string]string) error {
	verbose, _ := strconv.ParseBool(options["verbose"])
	language := languages[block.language]
	if language == "" {
		language = block.language
	}

	switch language {
	case "shell":
		if verbose {
			fmt.Printf("Executing code block\n\n%s\n", block)
		}
		// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
		args := []string{"-c", block.GetContent()}
		env := os.Environ()
		if len(options["echo"]) > 0 {
			args = append([]string{"-x"}, args...)
			// @see `-x` on https://www.gnu.org/software/bash/manual/html_node/The-Set-Builtin.html
			env = append(env, fmt.Sprintf("PS4=%s", options["echo"]))
		}
		cmd := exec.Command("sh", args...)
		cmd.Env = env

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			log.Fatal(err)
		}

		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		stdoutScanner := bufio.NewScanner(stdout)
		go func() {
			for stdoutScanner.Scan() {
				m := stdoutScanner.Text()
				fmt.Println(m)
			}
		}()

		stderrScanner := bufio.NewScanner(stderr)
		go func() {
			for stderrScanner.Scan() {
				m := stderrScanner.Text()
				fmt.Println(m)
			}
		}()

		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}

		return nil

	default:
		return fmt.Errorf("cannot handle language %s", language)
	}
}
