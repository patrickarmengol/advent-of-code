package assert

import (
	"strings"
	"testing"
)

func FileNotEmpty(t *testing.T, filename string, text string) {
	t.Helper()

	if len(strings.TrimSpace(text)) == 0 {
		t.Fatalf("%s file is empty", filename)
	}
}

func NilError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("encountered error %v", err)
	}
}

func Equal[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got: %v ; want %v", got, want)
	}
}
