package day08

import (
	_ "embed"
	"testing"

	"github.com/patrickarmengol/advent-of-code/2023/helpers/assert"
)

//go:embed input.txt
var inputText string

//go:embed example1.txt
var example1Text string

//go:embed example2.txt
var example2Text string

func TestPart1Example(t *testing.T) {
	assert.FileNotEmpty(t, "example1.txt", example1Text)

	expected := 6
	result, err := Part1(example1Text)

	assert.NilError(t, err)
	assert.Equal(t, result, expected)
}

func TestPart1Actual(t *testing.T) {
	assert.FileNotEmpty(t, "input.txt", inputText)

	expected := 17287
	result, err := Part1(inputText)

	assert.NilError(t, err)
	assert.Equal(t, result, expected)
}

func TestPart2Example(t *testing.T) {
	assert.FileNotEmpty(t, "example2.txt", example2Text)

	expected := 6
	result, err := Part2(example1Text)

	assert.NilError(t, err)
	assert.Equal(t, result, expected)
}

func TestPart2Actual(t *testing.T) {
	assert.FileNotEmpty(t, "input.txt", inputText)

	expected := 18625484023687
	result, err := Part2(inputText)

	assert.NilError(t, err)
	assert.Equal(t, result, expected)
}
