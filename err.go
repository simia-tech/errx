package errx

import (
	"fmt"
	"reflect"
	"runtime"
)

// Err defines the extended error struct.
type Err struct {
	message string

	cause error

	previous error

	file string
	line int
}

func NewErr(format string, args ...interface{}) Err {
	return Err{
		message: fmt.Sprintf(format, args...),
	}
}

func NewErrWithCause(other error, format string, args ...interface{}) Err {
	return Err{
		message:  fmt.Sprintf(format, args...),
		cause:    Cause(other),
		previous: other,
	}
}

func (e *Err) Location() (filename string, line int) {
	return e.file, e.line
}

func (e *Err) Underlying() error {
	return e.previous
}

func (e *Err) Cause() error {
	return e.cause
}

func (e *Err) Message() string {
	return e.message
}

func (e *Err) Error() string {
	err := e.previous
	if !sameError(Cause(err), e.cause) && e.cause != nil {
		err = e.cause
	}
	switch {
	case err == nil:
		return e.message
	case e.message == "":
		return err.Error()
	}
	return fmt.Sprintf("%s: %v", e.message, err)
}

func (e *Err) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case s.Flag('+'):
			fmt.Fprintf(s, "%s", ErrorStack(e))
			return
		case s.Flag('#'):
			// avoid infinite recursion by wrapping e into a type
			// that doesn't implement Formatter.
			fmt.Fprintf(s, "%#v", (*unformatter)(e))
			return
		}
		fallthrough
	case 's':
		fmt.Fprintf(s, "%s", e.Error())
	}
}

func (e *Err) SetLocation(callDepth int) {
	_, file, line, _ := runtime.Caller(callDepth + 1)
	e.file = trimGoPath(file)
	e.line = line
}

func sameError(e1, e2 error) bool {
	return reflect.DeepEqual(e1, e2)
}

type unformatter Err

func (unformatter) Format() { /* break the fmt.Formatter interface */ }
