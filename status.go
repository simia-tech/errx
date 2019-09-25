package errx

import (
	"fmt"
	"io"
	"net/http"
)

// All status errors for the 4xx and 5xx http status codes.
var (
	ErrBadRequest                   = Status(400)
	ErrUnauthorized                 = Status(401)
	ErrPaymentRequired              = Status(402)
	ErrForbidden                    = Status(403)
	ErrNotFound                     = Status(404)
	ErrMethodNotAllowed             = Status(405)
	ErrNotAcceptable                = Status(406)
	ErrProxyAuthRequired            = Status(407)
	ErrRequestTimeout               = Status(408)
	ErrConflict                     = Status(409)
	ErrGone                         = Status(410)
	ErrLengthRequired               = Status(411)
	ErrPreconditionFailed           = Status(412)
	ErrRequestEntityTooLarge        = Status(413)
	ErrRequestURITooLong            = Status(414)
	ErrUnsupportedMediaType         = Status(415)
	ErrRequestedRangeNotSatisfiable = Status(416)
	ErrExpectationFailed            = Status(417)
	ErrTeapot                       = Status(418)
	ErrMisdirectedRequest           = Status(421)
	ErrUnprocessableEntity          = Status(422)
	ErrLocked                       = Status(423)
	ErrFailedDependency             = Status(424)
	ErrTooEarly                     = Status(425)
	ErrUpgradeRequired              = Status(426)
	ErrPreconditionRequired         = Status(428)
	ErrTooManyRequests              = Status(429)
	ErrRequestHeaderFieldsTooLarge  = Status(431)
	ErrUnavailableForLegalReasons   = Status(451)

	ErrInternalServerError           = Status(500)
	ErrNotImplemented                = Status(501)
	ErrBadGateway                    = Status(502)
	ErrServiceUnavailable            = Status(503)
	ErrGatewayTimeout                = Status(504)
	ErrHTTPVersionNotSupported       = Status(505)
	ErrVariantAlsoNegotiates         = Status(506)
	ErrInsufficientStorage           = Status(507)
	ErrLoopDetected                  = Status(508)
	ErrNotExtended                   = Status(510)
	ErrNetworkAuthenticationRequired = Status(511)
)

// Status implements a status error which can be identified by a status code.
type Status int

// WriteText writes code and text to a http response as plain text.
func (s Status) WriteText(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(int(s))
	_, err := io.WriteString(w, s.String())
	return err
}

// WriteJSON writes code and text to a http response as json.
func (s Status) WriteJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(s))
	_, err := fmt.Fprintf(w, `{"error":%q}`, s.String())
	return err
}

func (s Status) String() string {
	return http.StatusText(int(s))
}

func (s Status) Error() string {
	return s.String()
}
