package errx

import (
	"errors"
)

// NotFoundf returns an error which satisfies IsNotFound().
func NotFoundf(format string, args ...interface{}) error {
	return Annotatef(ErrNotFound, format, args...)
}

// IsNotFound reports whether err was created with NotFoundf().
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// Unauthorizedf returns an error which satisfies IsUnauthorized().
func Unauthorizedf(format string, args ...interface{}) error {
	return Annotatef(ErrUnauthorized, format, args...)
}

// IsUnauthorized reports whether err was created with Unauthorizedf().
func IsUnauthorized(err error) bool {
	return errors.Is(err, ErrUnauthorized)
}

// NotImplementedf returns an error which satisfies IsNotImplemented().
func NotImplementedf(format string, args ...interface{}) error {
	return Annotatef(ErrNotImplemented, format, args...)
}

// IsNotImplemented reports whether err was created with
// NotImplementedf().
func IsNotImplemented(err error) bool {
	return errors.Is(err, ErrNotImplemented)
}

// AlreadyExistsf returns an error which satisfies IsAlreadyExists().
func AlreadyExistsf(format string, args ...interface{}) error {
	return Annotatef(ErrConflict, format, args...)
}

// IsAlreadyExists reports whether the error was created with
// AlreadyExistsf().
func IsAlreadyExists(err error) bool {
	return errors.Is(err, ErrConflict)
}

// Forbiddenf returns an error which satisfies IsForbidden().
func Forbiddenf(format string, args ...interface{}) error {
	return Annotatef(ErrForbidden, format, args...)
}

// IsForbidden reports whether the error was created with
// Forbiddenf().
func IsForbidden(err error) bool {
	return errors.Is(err, ErrForbidden)
}

// BadRequestf returns an error which satisfies IsBadRequest().
func BadRequestf(format string, args ...interface{}) error {
	return Annotatef(ErrBadRequest, format, args...)
}

// IsBadRequest reports whether err was created with BadRequestf().
func IsBadRequest(err error) bool {
	return errors.Is(err, ErrBadRequest)
}

// Timeoutf returns an error which satisfies IsTimeout().
func Timeoutf(format string, args ...interface{}) error {
	return Annotatef(ErrGatewayTimeout, format, args...)
}

// IsTimeout reports whether err was created with Timeoutf().
func IsTimeout(err error) bool {
	return errors.Is(err, ErrGatewayTimeout)
}
