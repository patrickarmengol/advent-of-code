package day03

import (
	_ "embed"
	"testing"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/assert"
)

//go:embed input.txt
var inputText string

//go:embed example.txt
var exampleText string

func TestPart1Example(t *testing.T) {
	assert.FileNotEmpty(t, "example.txt", exampleText)

	expected := 4361
	result, err := Part1(exampleText)

	assert.NilError(t, err)
	assert.Equal(t, result, expected)
}

func TestPart1Actual(t *testing.T) {
	assert.FileNotEmpty(t, "input.txt", inputText)

	expected := 530849
	result, err := Part1(inputText)

	assert.NilError(t, err)
	assert.Equal(t, result, expected)
}

func TestPart2Example(t *testing.T) {
	assert.FileNotEmpty(t, "example.txt", exampleText)

	expected := 467835
	result, err := Part2(exampleText)

	assert.NilError(t, err)
	assert.Equal(t, result, expected)
}

func TestPart2Actual(t *testing.T) {
	assert.FileNotEmpty(t, "input.txt", inputText)

	expected := 84900879
	result, err := Part2(inputText)

	assert.NilError(t, err)
	assert.Equal(t, result, expected)
}
