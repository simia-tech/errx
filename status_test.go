package errx_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/simia-tech/errx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatusWriteText(t *testing.T) {
	testFn := func(code int, expectText string) func(*testing.T) {
		return func(t *testing.T) {
			statusErr := errx.Status(code)

			rr := httptest.NewRecorder()
			require.NoError(t, statusErr.WriteText(rr))

			assert.Equal(t, code, rr.Code)
			assert.Equal(t, []string{"text/plain"}, rr.HeaderMap["Content-Type"])
			assert.Equal(t, expectText, rr.Body.String())
		}
	}

	t.Run("NotFound", testFn(http.StatusNotFound, "Not Found"))
	t.Run("Unauthorized", testFn(http.StatusUnauthorized, "Unauthorized"))
}

func TestStatusWriteJSON(t *testing.T) {
	testFn := func(code int, expectText string) func(*testing.T) {
		return func(t *testing.T) {
			statusErr := errx.Status(code)

			rr := httptest.NewRecorder()
			require.NoError(t, statusErr.WriteJSON(rr))

			assert.Equal(t, code, rr.Code)
			assert.Equal(t, []string{"application/json"}, rr.HeaderMap["Content-Type"])
			assert.Equal(t, expectText, rr.Body.String())
		}
	}

	t.Run("NotFound", testFn(http.StatusNotFound, `{"error":"Not Found"}`))
	t.Run("Unauthorized", testFn(http.StatusUnauthorized, `{"error":"Unauthorized"}`))
}
