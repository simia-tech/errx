package errx_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/simia-tech/errx"
)

func TestAnnotatef(t *testing.T) {
	err := errx.Annotatef(errors.New("three"), "one %s", "two")
	assert.Equal(t, "one two: three", err.Error())
}

func TestEqual(t *testing.T) {
	testFn := func(err, target error, expectResult bool) func(*testing.T) {
		return func(t *testing.T) {
			assert.Equal(t, expectResult, errors.Is(err, target))
		}
	}

	t.Run("Nil", testFn(nil, nil, true))
	t.Run("TwoAlreadyExists", testFn(errx.AlreadyExistsf("test"), errx.ErrConflict, true))
	t.Run("AlreadyExistsAndNotFound", testFn(errx.AlreadyExistsf("test"), errx.ErrNotFound, false))
	t.Run("AlreadyExistsAndNil", testFn(errx.AlreadyExistsf("test"), nil, false))
}
