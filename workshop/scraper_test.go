package workshop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetResourceTitle(t *testing.T) {
	title, err := GetResourceTitle(`https://steamcommunity.com/sharedfiles/filedetails/?id=2023507013`)
	assert.Nil(t, err)
	assert.Equal(t, "Vanilla Expanded Framework", title)
}
