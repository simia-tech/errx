package errx

import "fmt"

type notFound struct {
	Err
}

// NotFoundf returns an error which satisfies IsNotFound().
func NotFoundf(format string, args ...interface{}) error {
	return &notFound{wrap(nil, format, "", args...)}
}

// IsNotFound reports whether err was created with NotFoundf().
func IsNotFound(err error) bool {
	err = Cause(err)
	_, ok := err.(*notFound)
	return ok
}

type unauthorized struct {
	Err
}

// Unauthorizedf returns an error which satisfies IsUnauthorized().
func Unauthorizedf(format string, args ...interface{}) error {
	return &unauthorized{wrap(nil, format, "", args...)}
}

// IsUnauthorized reports whether err was created with Unauthorizedf().
func IsUnauthorized(err error) bool {
	err = Cause(err)
	_, ok := err.(*unauthorized)
	return ok
}

type notImplemented struct {
	Err
}

// NotImplementedf returns an error which satisfies IsNotImplemented().
func NotImplementedf(format string, args ...interface{}) error {
	return &notImplemented{wrap(nil, format, "", args...)}
}

// IsNotImplemented reports whether err was created with
// NotImplementedf().
func IsNotImplemented(err error) bool {
	err = Cause(err)
	_, ok := err.(*notImplemented)
	return ok
}

type alreadyExists struct {
	Err
}

// AlreadyExistsf returns an error which satisfies IsAlreadyExists().
func AlreadyExistsf(format string, args ...interface{}) error {
	return &alreadyExists{wrap(nil, format, "", args...)}
}

// IsAlreadyExists reports whether the error was created with
// AlreadyExistsf().
func IsAlreadyExists(err error) bool {
	err = Cause(err)
	_, ok := err.(*alreadyExists)
	return ok
}

type badRequest struct {
	Err
}

// BadRequestf returns an error which satisfies IsBadRequest().
func BadRequestf(format string, args ...interface{}) error {
	return &badRequest{wrap(nil, format, "", args...)}
}

// IsBadRequest reports whether err was created with BadRequestf().
func IsBadRequest(err error) bool {
	err = Cause(err)
	_, ok := err.(*badRequest)
	return ok
}

type timeout struct {
	Err
}

// Timeoutf returns an error which satisfies IsTimeout().
func Timeoutf(format string, args ...interface{}) error {
	return &timeout{wrap(nil, format, "", args...)}
}

// IsTimeout reports whether err was created with Timeoutf().
func IsTimeout(err error) bool {
	err = Cause(err)
	_, ok := err.(*timeout)
	return ok
}

func wrap(err error, format, suffix string, args ...interface{}) Err {
	newErr := Err{
		message:  fmt.Sprintf(format+suffix, args...),
		previous: err,
	}
	newErr.SetLocation(2)
	return newErr
}
