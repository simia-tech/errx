package errx

import (
	"log"
	"net/http"
)

// WriteHTTPStatus writes the provided error to the provided response writer.
func WriteHTTPStatus(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	statusCode := http.StatusInternalServerError

	switch {
	case IsAlreadyExists(err):
		statusCode = http.StatusConflict
	case IsBadRequest(err):
		statusCode = http.StatusBadRequest
	case IsNotFound(err):
		statusCode = http.StatusNotFound
	case IsNotImplemented(err):
		statusCode = http.StatusNotImplemented
	case IsTimeout(err):
		statusCode = http.StatusBadGateway
	case IsUnauthorized(err):
		statusCode = http.StatusUnauthorized
	case IsForbidden(err):
		statusCode = http.StatusForbidden
	}

	if statusCode >= 400 && statusCode != http.StatusNotFound {
		log.Printf("error: %v", err)
	}

	w.WriteHeader(statusCode)
}
