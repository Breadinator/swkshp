package workshop_tests

import (
	"testing"

	"github.com/breadinator/swkshp/workshop"
	"github.com/stretchr/testify/assert"
)

func Test_GetResourceTitle(t *testing.T) {
	title, err := workshop.GetResourceTitle(`https://steamcommunity.com/sharedfiles/filedetails/?id=2023507013`)
	assert.Nil(t, err)
	assert.Equal(t, "Vanilla Expanded Framework", title)
}
