package errx

import (
	"fmt"
	"strings"
)

func ErrorStack(err error) string {
	return strings.Join(errorStack(err), "\n")
}

func errorStack(err error) []string {
	if err == nil {
		return nil
	}

	// We want the first error first
	var lines []string
	for {
		var buff []byte
		if err, ok := err.(locationer); ok {
			file, line := err.Location()
			// Strip off the leading GOPATH/src path elements.
			file = trimGoPath(file)
			if file != "" {
				buff = append(buff, fmt.Sprintf("%s:%d", file, line)...)
				buff = append(buff, ": "...)
			}
		}
		if cerr, ok := err.(wrapper); ok {
			message := cerr.Message()
			buff = append(buff, message...)
			// If there is a cause for this error, and it is different to the cause
			// of the underlying error, then output the error string in the stack trace.
			var cause error
			if err1, ok := err.(causer); ok {
				cause = err1.Cause()
			}
			err = cerr.Underlying()
			if cause != nil && !sameError(Cause(err), cause) {
				if message != "" {
					buff = append(buff, ": "...)
				}
				buff = append(buff, cause.Error()...)
			}
		} else {
			buff = append(buff, err.Error()...)
			err = nil
		}
		lines = append(lines, string(buff))
		if err == nil {
			break
		}
	}
	// reverse the lines to get the original error, which was at the end of
	// the list, back to the start.
	var result []string
	for i := len(lines); i > 0; i-- {
		result = append(result, lines[i-1])
	}
	return result
}
