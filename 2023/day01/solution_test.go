package day01

import (
	_ "embed"
	"testing"

	"github.com/patrickarmengol/advent-of-code/2023/helpers"
)

//go:embed input.txt
var inputText string

//go:embed example1.txt
var example1Text string

//go:embed example2.txt
var example2Text string

func TestExamplePart1(t *testing.T) {
	if helpers.IsEmpty(example1Text) {
		t.Fatalf("empty example.txt file")
	}

	expected := "142"
	result, err := Part1(example1Text)
	if err != nil {
		t.Errorf("encountered error: %v", err)
	}
	if expected != result {
		t.Errorf("got: %v; want %v", result, expected)
	}
}

func TestActualPart1(t *testing.T) {
	if helpers.IsEmpty(inputText) {
		t.Fatalf("empty input.txt file")
	}

	expected := "53921"
	result, err := Part1(inputText)
	if err != nil {
		t.Errorf("encountered error: %v", err)
	}
	if expected != result {
		t.Errorf("got: %v; want %v", result, expected)
	}
}

func TestExamplePart2(t *testing.T) {
	if helpers.IsEmpty(example2Text) {
		t.Fatalf("empty example.txt file")
	}

	expected := "281"
	result, err := Part2(example2Text)
	if err != nil {
		t.Errorf("encountered error: %v", err)
	}
	if expected != result {
		t.Errorf("got: %v; want %v", result, expected)
	}
}

func TestActualPart2(t *testing.T) {
	if helpers.IsEmpty(inputText) {
		t.Fatalf("empty input.txt file")
	}

	expected := "54676"
	result, err := Part2(inputText)
	if err != nil {
		t.Errorf("encountered error: %v", err)
	}
	if expected != result {
		t.Errorf("got: %v; want %v", result, expected)
	}
}
