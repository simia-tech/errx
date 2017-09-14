package errx

import (
	"fmt"
	"reflect"
)

func Errorf(format string, args ...interface{}) error {
	err := &Err{message: fmt.Sprintf(format, args...)}
	err.SetLocation(1)
	return err
}

func Annotatef(other error, format string, args ...interface{}) error {
	if other == nil {
		return nil
	}
	err := &Err{
		previous: other,
		cause:    Cause(other),
		message:  fmt.Sprintf(format, args...),
	}
	err.SetLocation(1)
	return err
}

func IsErr(err error) bool {
	if IsAlreadyExists(err) || IsBadRequest(err) || IsNotFound(err) || IsNotImplemented(err) || IsTimeout(err) || IsUnauthorized(err) {
		return true
	}
	_, ok := err.(*Err)
	return ok
}

func Equal(a, b error) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	at, bt := reflect.TypeOf(a), reflect.TypeOf(b)
	if at != nil && bt != nil && at.String() == bt.String() {
		return a.Error() == b.Error()
	}
	return false
}
