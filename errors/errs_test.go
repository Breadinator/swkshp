package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Wrap(t *testing.T) {
	assert.True(t, Is(Wrap(ErrHTTP, "%d", 404), ErrHTTP))
	assert.False(t, Is(Wrap(ErrHTTP, "%d", 404), ErrGameNotFound))
}
