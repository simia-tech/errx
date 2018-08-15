package errx

import "testing"

// AssertError is an assert helper to check for an expected error.
func AssertError(tb testing.TB, expected, actual error) {
	if !Equal(expected, actual) {
		tb.Errorf("expected error [%T] [%v], got [%T] [%v]", expected, expected, actual, actual)
	}
}

// RequireError is a require helper to check for an expected error.
func RequireError(tb testing.TB, expected, actual error) {
	if !Equal(expected, actual) {
		tb.Fatalf("expected error [%T] [%v], got [%T] [%v]", expected, expected, actual, actual)
	}
}
