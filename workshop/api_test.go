package workshop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_A(t *testing.T) {
	_, a := CheckAvailable(`https://steamcommunity.com/sharedfiles/filedetails/?id=2683996590`)
	assert.Nil(t, a)

	_, b := CheckAvailable(`https://steamcommunity.com/sharedfiles/filedetails/?id=123`)
	assert.Nil(t, b)
}
