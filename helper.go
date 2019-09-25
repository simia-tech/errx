package errx

import (
	"fmt"
)

// Annotatef annotates the provided error with the provided text and arguments.
func Annotatef(other error, format string, args ...interface{}) error {
	if other == nil {
		return nil
	}
	args = append(args, other)
	return fmt.Errorf(format+": %w", args...)
}
