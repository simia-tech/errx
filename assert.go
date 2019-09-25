package errx

import (
	"errors"
	"testing"
)

// AssertError is an assert helper to check for an expected error.
func AssertError(tb testing.TB, expected, actual error) {
	if !errors.Is(actual, expected) {
		tb.Errorf("expected error [%T] [%v], got [%T] [%v]", expected, expected, actual, actual)
	}
}

// RequireError is a require helper to check for an expected error.
func RequireError(tb testing.TB, expected, actual error) {
	if !errors.Is(actual, expected) {
		tb.Fatalf("expected error [%T] [%v], got [%T] [%v]", expected, expected, actual, actual)
	}
}
