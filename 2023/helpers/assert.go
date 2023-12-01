package helpers

import "strings"

func IsEmpty(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}
