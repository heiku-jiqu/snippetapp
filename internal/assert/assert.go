package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func StringContains(t *testing.T, actual, substr string) {
	t.Helper()
	if !strings.Contains(actual, substr) {
		t.Errorf("expected %q to contain %q", actual, substr)
	}
}
