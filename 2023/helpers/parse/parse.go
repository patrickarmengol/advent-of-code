package parse

import (
	"strings"
)

func Lines(text string) []string {
	return strings.Split(strings.TrimRight(text, "\n"), "\n")
}
