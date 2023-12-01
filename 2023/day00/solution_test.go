package day00

import (
	_ "embed"
	"testing"

	"github.com/patrickarmengol/advent-of-code/2023/helpers"
)

func TestSolutions(t *testing.T) {
	// check text files populated
	if helpers.IsEmpty(exampleText) {
		t.Fatalf("empty example.txt file")
	}
	if helpers.IsEmpty(inputText) {
		t.Fatalf("empty input.txt file")
	}

	// setup tests (expected must be filled while solving)
	tests := []struct {
		name     string
		data     string
		sol      func(string) string
		expected string
	}{
		{
			name:     "example-part1",
			data:     exampleText,
			sol:      Part1,
			expected: "__UNKNOWN__",
		},
		{
			name:     "actual-part1",
			data:     inputText,
			sol:      Part1,
			expected: "__UNKNOWN__",
		},
		{
			name:     "example-part2",
			data:     exampleText,
			sol:      Part2,
			expected: "__UNKNOWN__",
		},
		{
			name:     "actual-part2",
			data:     inputText,
			sol:      Part2,
			expected: "__UNKNOWN__",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.sol(tt.data)
			if result != tt.expected {
				t.Errorf("got: %v; want %v", result, tt.expected)
			}
		})
	}
}
