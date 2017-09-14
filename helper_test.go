package errx_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"code.posteo.de/common/errx"
)

func TestErrorf(t *testing.T) {
	err := errx.Errorf("one %s", "two")
	assert.Equal(t, "one two", err.Error())
}

func TestAnnotatef(t *testing.T) {
	err := errx.Annotatef(errors.New("three"), "one %s", "two")
	assert.Equal(t, "one two: three", err.Error())
}

func TestEqual(t *testing.T) {
	testCases := []struct {
		name   string
		a      error
		b      error
		result bool
	}{
		{"Nil", nil, nil, true},
		{"TwoAlreadyExists", errx.AlreadyExistsf("test"), errx.AlreadyExistsf("test"), true},
		{"AlreadyExistsAndNotFound", errx.AlreadyExistsf("test"), errx.NotFoundf("test"), false},
		{"AlreadyExistsAndNil", errx.AlreadyExistsf("test"), nil, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.result, errx.Equal(testCase.a, testCase.b))
		})
	}
}
