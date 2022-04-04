package workshop_tests

import (
	"testing"

	"github.com/breadinator/swkshp/workshop"
	"github.com/stretchr/testify/assert"
)

func Test_GetGame_UsingString(t *testing.T) {
	url := `https://steamcommunity.com/sharedfiles/filedetails/?id=818773962`
	game, err := workshop.GetGame(url)
	assert.Nil(t, err)
	if err == nil {
		assert.Equal(t, "RimWorld", game)
	}
}

func Test_GetGame_UsingInt(t *testing.T) {
	id := 818773962
	game, err := workshop.GetGame(id)
	assert.Nil(t, err)
	if err == nil {
		assert.Equal(t, "RimWorld", game)
	}
}

func Test_GetGame_UsingInvalidType(t *testing.T) {
	check := func(a interface{}) {
		_, err := workshop.GetGame(a)
		assert.NotNil(t, err)
	}

	check(1.2)
	check(nil)
	check(false)
	check('c')
}
