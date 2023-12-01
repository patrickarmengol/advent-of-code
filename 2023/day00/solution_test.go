package day00

import (
	_ "embed"
	"testing"

	"github.com/patrickarmengol/advent-of-code/2023/helpers"
)

//go:embed input.txt
var inputText string

//go:embed example.txt
var exampleText string

func TestPart1Example(t *testing.T) {
	if helpers.IsEmpty(exampleText) {
		t.Fatalf("empty example.txt file")
	}

	expected := "__UNKNOWN__"
	result, err := Part1(exampleText)
	if err != nil {
		t.Errorf("encountered error: %v", err)
	}
	if expected != result {
		t.Errorf("got: %v; want %v", result, expected)
	}
}

func TestPart1Actual(t *testing.T) {
	if helpers.IsEmpty(inputText) {
		t.Fatalf("empty input.txt file")
	}

	expected := "__UNKNOWN__"
	result, err := Part1(inputText)
	if err != nil {
		t.Errorf("encountered error: %v", err)
	}
	if expected != result {
		t.Errorf("got: %v; want %v", result, expected)
	}
}

func TestPart2Example(t *testing.T) {
	if helpers.IsEmpty(exampleText) {
		t.Fatalf("empty example.txt file")
	}

	expected := "__UNKNOWN__"
	result, err := Part2(exampleText)
	if err != nil {
		t.Errorf("encountered error: %v", err)
	}
	if expected != result {
		t.Errorf("got: %v; want %v", result, expected)
	}
}

func TestPart2Actual(t *testing.T) {
	if helpers.IsEmpty(inputText) {
		t.Fatalf("empty input.txt file")
	}

	expected := "__UNKNOWN__"
	result, err := Part2(inputText)
	if err != nil {
		t.Errorf("encountered error: %v", err)
	}
	if expected != result {
		t.Errorf("got: %v; want %v", result, expected)
	}
}
