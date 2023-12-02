package parse

import (
	"strings"
)

func Lines(text string) []string {
	return strings.Split(text, "\n")
}

func Words(text string) []string {
	return strings.Split(text, " ")
}
