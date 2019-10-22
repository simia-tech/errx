package errx

import (
	"errors"
	"testing"
)

// AssertError is an assert helper to check for an expected error.
func AssertError(tb testing.TB, expected, actual error) {
	if expectedStatusErr, ok := expected.(*Status); ok {
		var actualStatusErr *Status
		if errors.As(actual, &actualStatusErr) {
			if expectedStatusErr.Code != actualStatusErr.Code {
				tb.Errorf("expected error status code %d, got %d", expectedStatusErr.Code, actualStatusErr.Code)
			}
			if expectedStatusErr.Text != actualStatusErr.Text {
				tb.Errorf("expected error text [%s], got [%s]", expectedStatusErr.Text, actualStatusErr.Text)
			}
			return
		}
	}
	if !errors.Is(actual, expected) {
		tb.Errorf("expected error [%T / %v], got [%T / %v]", expected, expected, actual, actual)
	}
}

// RequireError is a require helper to check for an expected error.
func RequireError(tb testing.TB, expected, actual error) {
	if expectedStatusErr, ok := expected.(*Status); ok {
		var actualStatusErr *Status
		if errors.As(actual, &actualStatusErr) {
			if expectedStatusErr.Code != actualStatusErr.Code {
				tb.Fatalf("expected error status code %d, got %d", expectedStatusErr.Code, actualStatusErr.Code)
			}
			if expectedStatusErr.Text != actualStatusErr.Text {
				tb.Fatalf("expected error text [%s], got [%s]", expectedStatusErr.Text, actualStatusErr.Text)
			}
			return
		}
	}
	if !errors.Is(actual, expected) {
		tb.Fatalf("expected error [%T / %v], got [%T / %v]", expected, expected, actual, actual)
	}
}
