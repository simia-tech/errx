package errx

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

// All status errors for the 4xx and 5xx http status codes.
var (
	ErrBadRequest                   = &Status{Code: http.StatusBadRequest}
	ErrUnauthorized                 = &Status{Code: http.StatusUnauthorized}
	ErrPaymentRequired              = &Status{Code: http.StatusPaymentRequired}
	ErrForbidden                    = &Status{Code: http.StatusForbidden}
	ErrNotFound                     = &Status{Code: http.StatusNotFound}
	ErrMethodNotAllowed             = &Status{Code: http.StatusMethodNotAllowed}
	ErrNotAcceptable                = &Status{Code: http.StatusNotAcceptable}
	ErrProxyAuthRequired            = &Status{Code: http.StatusProxyAuthRequired}
	ErrRequestTimeout               = &Status{Code: http.StatusRequestTimeout}
	ErrConflict                     = &Status{Code: http.StatusConflict}
	ErrGone                         = &Status{Code: http.StatusGone}
	ErrLengthRequired               = &Status{Code: http.StatusLengthRequired}
	ErrPreconditionFailed           = &Status{Code: http.StatusPreconditionFailed}
	ErrRequestEntityTooLarge        = &Status{Code: http.StatusRequestEntityTooLarge}
	ErrRequestURITooLong            = &Status{Code: http.StatusRequestURITooLong}
	ErrUnsupportedMediaType         = &Status{Code: http.StatusUnsupportedMediaType}
	ErrRequestedRangeNotSatisfiable = &Status{Code: http.StatusRequestedRangeNotSatisfiable}
	ErrExpectationFailed            = &Status{Code: http.StatusExpectationFailed}
	ErrTeapot                       = &Status{Code: http.StatusTeapot}
	ErrMisdirectedRequest           = &Status{Code: http.StatusMisdirectedRequest}
	ErrUnprocessableEntity          = &Status{Code: http.StatusUnprocessableEntity}
	ErrLocked                       = &Status{Code: http.StatusLocked}
	ErrFailedDependency             = &Status{Code: http.StatusFailedDependency}
	ErrTooEarly                     = &Status{Code: http.StatusTooEarly}
	ErrUpgradeRequired              = &Status{Code: http.StatusUpgradeRequired}
	ErrPreconditionRequired         = &Status{Code: http.StatusPreconditionRequired}
	ErrTooManyRequests              = &Status{Code: http.StatusTooManyRequests}
	ErrRequestHeaderFieldsTooLarge  = &Status{Code: http.StatusRequestHeaderFieldsTooLarge}
	ErrUnavailableForLegalReasons   = &Status{Code: http.StatusUnavailableForLegalReasons}

	ErrInternalServerError           = &Status{Code: http.StatusInternalServerError}
	ErrNotImplemented                = &Status{Code: http.StatusNotImplemented}
	ErrBadGateway                    = &Status{Code: http.StatusBadGateway}
	ErrServiceUnavailable            = &Status{Code: http.StatusServiceUnavailable}
	ErrGatewayTimeout                = &Status{Code: http.StatusGatewayTimeout}
	ErrHTTPVersionNotSupported       = &Status{Code: http.StatusHTTPVersionNotSupported}
	ErrVariantAlsoNegotiates         = &Status{Code: http.StatusVariantAlsoNegotiates}
	ErrInsufficientStorage           = &Status{Code: http.StatusInsufficientStorage}
	ErrLoopDetected                  = &Status{Code: http.StatusLoopDetected}
	ErrNotExtended                   = &Status{Code: http.StatusNotExtended}
	ErrNetworkAuthenticationRequired = &Status{Code: http.StatusNetworkAuthenticationRequired}
)

// Status implements a status error which can be identified by a status code.
type Status struct {
	Code int
	Text string
}

// StatusErr converts the provided error into a status error.
func StatusErr(err error) *Status {
	if err == nil {
		return nil
	}

	statusErr := (*Status)(nil)
	if errors.As(err, &statusErr) {
		statusErr.Text = err.Error()
		return statusErr
	}

	return &Status{Code: http.StatusInternalServerError, Text: err.Error()}
}

// Write writes code and text to the provided http response. The body format is determind by
// the `Content-Type` header in the `http.ResponseWriter`.
func (s *Status) Write(w http.ResponseWriter) error {
	w.WriteHeader(s.Code)
	switch ct := w.Header().Get("Content-Type"); ct {
	case "application/json":
		_, err := fmt.Fprintf(w, `{"error":%q}`, s.String())
		return err
	default:
		_, err := io.WriteString(w, s.String())
		return err
	}
}

// WriteText writes code and text to a http response as plain text.
func (s *Status) WriteText(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(s.Code)
	_, err := io.WriteString(w, s.String())
	return err
}

// WriteJSON writes code and text to a http response as json.
func (s Status) WriteJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.Code)
	_, err := fmt.Fprintf(w, `{"error":%q}`, s.String())
	return err
}

func (s Status) String() string {
	if s.Text == "" {
		return http.StatusText(s.Code)
	}
	return s.Text
}

func (s Status) Error() string {
	return s.String()
}

// ErrBadRequestf returns a new status with the corresponding status code.
func ErrBadRequestf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusBadRequest, Text: fmt.Sprintf(format, arguments...)}
}

// ErrUnauthorizedf returns a new status with the corresponding status code.
func ErrUnauthorizedf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusUnauthorized, Text: fmt.Sprintf(format, arguments...)}
}

// ErrPaymentRequiredf returns a new status with the corresponding status code.
func ErrPaymentRequiredf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusPaymentRequired, Text: fmt.Sprintf(format, arguments...)}
}

// ErrForbiddenf returns a new status with the corresponding status code.
func ErrForbiddenf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusForbidden, Text: fmt.Sprintf(format, arguments...)}
}

// ErrNotFoundf returns a new status with the corresponding status code.
func ErrNotFoundf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusNotFound, Text: fmt.Sprintf(format, arguments...)}
}

// ErrMethodNotAllowedf returns a new status with the corresponding status code.
func ErrMethodNotAllowedf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusMethodNotAllowed, Text: fmt.Sprintf(format, arguments...)}
}

// ErrNotAcceptablef returns a new status with the corresponding status code.
func ErrNotAcceptablef(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusNotAcceptable, Text: fmt.Sprintf(format, arguments...)}
}

// ErrProxyAuthRequiredf returns a new status with the corresponding status code.
func ErrProxyAuthRequiredf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusProxyAuthRequired, Text: fmt.Sprintf(format, arguments...)}
}

// ErrRequestTimeoutf returns a new status with the corresponding status code.
func ErrRequestTimeoutf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusRequestTimeout, Text: fmt.Sprintf(format, arguments...)}
}

// ErrConflictf returns a new status with the corresponding status code.
func ErrConflictf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusConflict, Text: fmt.Sprintf(format, arguments...)}
}

// ErrGonef returns a new status with the corresponding status code.
func ErrGonef(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusGone, Text: fmt.Sprintf(format, arguments...)}
}

// ErrLengthRequiredf returns a new status with the corresponding status code.
func ErrLengthRequiredf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusLengthRequired, Text: fmt.Sprintf(format, arguments...)}
}

// ErrPreconditionFailedf returns a new status with the corresponding status code.
func ErrPreconditionFailedf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusPreconditionFailed, Text: fmt.Sprintf(format, arguments...)}
}

// ErrRequestEntityTooLargef returns a new status with the corresponding status code.
func ErrRequestEntityTooLargef(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusRequestEntityTooLarge, Text: fmt.Sprintf(format, arguments...)}
}

// ErrRequestURITooLongf returns a new status with the corresponding status code.
func ErrRequestURITooLongf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusRequestURITooLong, Text: fmt.Sprintf(format, arguments...)}
}

// ErrUnsupportedMediaTypef returns a new status with the corresponding status code.
func ErrUnsupportedMediaTypef(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusUnsupportedMediaType, Text: fmt.Sprintf(format, arguments...)}
}

// ErrRequestedRangeNotSatisfiablef returns a new status with the corresponding status code.
func ErrRequestedRangeNotSatisfiablef(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusRequestedRangeNotSatisfiable, Text: fmt.Sprintf(format, arguments...)}
}

// ErrExpectationFailedf returns a new status with the corresponding status code.
func ErrExpectationFailedf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusExpectationFailed, Text: fmt.Sprintf(format, arguments...)}
}

// ErrTeapotf returns a new status with the corresponding status code.
func ErrTeapotf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusTeapot, Text: fmt.Sprintf(format, arguments...)}
}

// ErrMisdirectedRequestf returns a new status with the corresponding status code.
func ErrMisdirectedRequestf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusMisdirectedRequest, Text: fmt.Sprintf(format, arguments...)}
}

// ErrUnprocessableEntityf returns a new status with the corresponding status code.
func ErrUnprocessableEntityf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusUnprocessableEntity, Text: fmt.Sprintf(format, arguments...)}
}

// ErrLockedf returns a new status with the corresponding status code.
func ErrLockedf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusLocked, Text: fmt.Sprintf(format, arguments...)}
}

// ErrFailedDependencyf returns a new status with the corresponding status code.
func ErrFailedDependencyf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusFailedDependency, Text: fmt.Sprintf(format, arguments...)}
}

// ErrTooEarlyf returns a new status with the corresponding status code.
func ErrTooEarlyf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusTooEarly, Text: fmt.Sprintf(format, arguments...)}
}

// ErrUpgradeRequiredf returns a new status with the corresponding status code.
func ErrUpgradeRequiredf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusUpgradeRequired, Text: fmt.Sprintf(format, arguments...)}
}

// ErrPreconditionRequiredf returns a new status with the corresponding status code.
func ErrPreconditionRequiredf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusPreconditionRequired, Text: fmt.Sprintf(format, arguments...)}
}

// ErrTooManyRequestsf returns a new status with the corresponding status code.
func ErrTooManyRequestsf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusTooManyRequests, Text: fmt.Sprintf(format, arguments...)}
}

// ErrRequestHeaderFieldsTooLargef returns a new status with the corresponding status code.
func ErrRequestHeaderFieldsTooLargef(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusRequestHeaderFieldsTooLarge, Text: fmt.Sprintf(format, arguments...)}
}

// ErrUnavailableForLegalReasonsf returns a new status with the corresponding status code.
func ErrUnavailableForLegalReasonsf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusUnavailableForLegalReasons, Text: fmt.Sprintf(format, arguments...)}
}

// ErrInternalServerErrorf returns a new status with the corresponding status code.
func ErrInternalServerErrorf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusInternalServerError, Text: fmt.Sprintf(format, arguments...)}
}

// ErrNotImplementedf returns a new status with the corresponding status code.
func ErrNotImplementedf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusNotImplemented, Text: fmt.Sprintf(format, arguments...)}
}

// ErrBadGatewayf returns a new status with the corresponding status code.
func ErrBadGatewayf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusBadGateway, Text: fmt.Sprintf(format, arguments...)}
}

// ErrServiceUnavailablef returns a new status with the corresponding status code.
func ErrServiceUnavailablef(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusServiceUnavailable, Text: fmt.Sprintf(format, arguments...)}
}

// ErrGatewayTimeoutf returns a new status with the corresponding status code.
func ErrGatewayTimeoutf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusGatewayTimeout, Text: fmt.Sprintf(format, arguments...)}
}

// ErrHTTPVersionNotSupportedf returns a new status with the corresponding status code.
func ErrHTTPVersionNotSupportedf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusHTTPVersionNotSupported, Text: fmt.Sprintf(format, arguments...)}
}

// ErrVariantAlsoNegotiatesf returns a new status with the corresponding status code.
func ErrVariantAlsoNegotiatesf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusVariantAlsoNegotiates, Text: fmt.Sprintf(format, arguments...)}
}

// ErrInsufficientStoragef returns a new status with the corresponding status code.
func ErrInsufficientStoragef(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusInsufficientStorage, Text: fmt.Sprintf(format, arguments...)}
}

// ErrLoopDetectedf returns a new status with the corresponding status code.
func ErrLoopDetectedf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusLoopDetected, Text: fmt.Sprintf(format, arguments...)}
}

// ErrNotExtendedf returns a new status with the corresponding status code.
func ErrNotExtendedf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusNotExtended, Text: fmt.Sprintf(format, arguments...)}
}

// ErrNetworkAuthenticationRequiredf returns a new status with the corresponding status code.
func ErrNetworkAuthenticationRequiredf(format string, arguments ...interface{}) *Status {
	return &Status{Code: http.StatusNetworkAuthenticationRequired, Text: fmt.Sprintf(format, arguments...)}
}
