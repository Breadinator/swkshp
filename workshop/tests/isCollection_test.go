package workshop_tests

import (
	"testing"

	"github.com/breadinator/swkshp/workshop"
	"github.com/stretchr/testify/assert"
)

func Test_IsCollection(t *testing.T) {
	check := func(url string, expected bool) {
		isCollection, err := workshop.IsCollection(url)
		assert.Nil(t, err)
		assert.Equal(t, isCollection, expected)
	}

	check("https://steamcommunity.com/workshop/filedetails/?id=1884025115", true)
	check("https://steamcommunity.com/workshop/filedetails/?id=818773962", false)
}
