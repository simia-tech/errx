package errx_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/simia-tech/errx/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatusErr(t *testing.T) {
	testFn := func(err error, expectCode int, expectString string) func(*testing.T) {
		return func(t *testing.T) {
			statusErr := errx.StatusErr(err)
			assert.Equal(t, expectCode, statusErr.Code)
			assert.Equal(t, expectString, statusErr.String())
		}
	}

	t.Run("DefaultError", testFn(fmt.Errorf("error"), http.StatusInternalServerError, "error"))
	t.Run("WrappedDefaultError", testFn(fmt.Errorf("wrap: %w", fmt.Errorf("error")), http.StatusInternalServerError, "wrap: error"))
	t.Run("HttpError", testFn(errx.ErrNotFoundf("not found"), http.StatusNotFound, "not found"))
	t.Run("WrappedHttpError", testFn(fmt.Errorf("wrap: %w", errx.ErrNotFoundf("not found")), http.StatusNotFound, "wrap: not found"))
}

func TestStatusWrite(t *testing.T) {
	testFn := func(code int, contentType, expectBody string) func(*testing.T) {
		return func(t *testing.T) {
			statusErr := &errx.Status{Code: code}

			rr := httptest.NewRecorder()
			rr.Header().Set("Content-Type", contentType)
			require.NoError(t, statusErr.Write(rr))

			assert.Equal(t, code, rr.Code)
			assert.Equal(t, expectBody, rr.Body.String())
		}
	}

	t.Run("NotFoundText", testFn(http.StatusNotFound, "text/plain", "Not Found"))
	t.Run("NotFoundJSON", testFn(http.StatusNotFound, "application/json", `{"error":"Not Found"}`))
}

func TestStatusWriteText(t *testing.T) {
	testFn := func(code int, text string, expectBody string) func(*testing.T) {
		return func(t *testing.T) {
			statusErr := &errx.Status{Code: code, Text: text}

			rr := httptest.NewRecorder()
			require.NoError(t, statusErr.WriteText(rr))

			assert.Equal(t, code, rr.Code)
			assert.Equal(t, []string{"text/plain"}, rr.HeaderMap["Content-Type"])
			assert.Equal(t, expectBody, rr.Body.String())
		}
	}

	t.Run("NotFound", testFn(http.StatusNotFound, "", "Not Found"))
	t.Run("NotFoundWithText", testFn(http.StatusNotFound, "not found", "not found"))
}

func TestStatusWriteJSON(t *testing.T) {
	testFn := func(code int, text string, expectBody string) func(*testing.T) {
		return func(t *testing.T) {
			statusErr := &errx.Status{Code: code, Text: text}

			rr := httptest.NewRecorder()
			require.NoError(t, statusErr.WriteJSON(rr))

			assert.Equal(t, code, rr.Code)
			assert.Equal(t, []string{"application/json"}, rr.HeaderMap["Content-Type"])
			assert.Equal(t, expectBody, rr.Body.String())
		}
	}

	t.Run("NotFound", testFn(http.StatusNotFound, "", `{"error":"Not Found"}`))
	t.Run("NotFoundWithText", testFn(http.StatusNotFound, "not found", `{"error":"not found"}`))
}
