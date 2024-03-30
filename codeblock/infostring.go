package codeblock

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type InfoString struct {
	name       string
	properties map[string]string
}

func (infoString InfoString) GetName() string {
	return infoString.name
}

func (infoString InfoString) GetProperty(name string) string {
	return infoString.properties[name]
}

var namePattern = regexp.MustCompile("(?i)^[a-z0-9-]+$")

func isValidName(name string) bool {
	return namePattern.MatchString(name)
}

// https://github.github.com/gfm/#info-string
func ParseInfoString(infoString string) (InfoString, error) {
	result := InfoString{}
	trimmed := strings.TrimSpace(infoString)
	if len(trimmed) == 0 {
		return result, nil
	}

	wrapperName := "infostring"
	input := fmt.Sprintf("<%[1]s><%[2]s></%[1]s>", wrapperName, trimmed)

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return result, err
	}

	err = fmt.Errorf("error parsing info string %#q", infoString)
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == wrapperName {
			n = n.FirstChild
			if isValidName(n.Data) {
				result.name = n.Data
				result.properties = map[string]string{}
				for _, a := range n.Attr {
					result.properties[a.Key] = a.Val
				}
				err = nil
			}
		}
		for c := n.FirstChild; c != nil && err != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	return result, err
}

func (infoString InfoString) String() string {
	return fmt.Sprintf("%s %q", infoString.name, infoString.properties)
}
