package helpers

import (
	"bufio"
	"fmt"
	"strings"
)

func GetLines(text string) ([]string, error) {
	var lines []string

	s := bufio.NewScanner(strings.NewReader(text))
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if s.Err() != nil {
		return nil, fmt.Errorf("failed to scan reader: %w", s.Err())
	}

	return lines, nil
}
