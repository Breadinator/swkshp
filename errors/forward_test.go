package errors

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: add more/better test cases
var tests = [...][2]error{
	{Wrap(ErrHTTP, "%d", 404), ErrHTTP},
	{ErrHTTP, ErrGameNotFound},
	{Wrap(os.ErrNotExist, "lol"), os.ErrNotExist},
	{nil, nil},
	{nil, ErrIllegalPath},
	{ErrIllegalPath, nil},
}

func Test_Is_As(t *testing.T) {
	for _, test := range tests {
		a, b := test[0], test[1]
		assert.Equal(t, Is(a, b), errors.Is(a, b))
		assert.Equal(t, As(a, &b), errors.As(a, &b))
		assert.Equal(t, Unwrap(a), errors.Unwrap(a))
		assert.Equal(t, Unwrap(b), errors.Unwrap(b))
	}
}
