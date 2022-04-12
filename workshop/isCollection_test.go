package workshop

import (
	"testing"

	"github.com/breadinator/swkshp/errors"
	"github.com/stretchr/testify/assert"
)

func Test_IsCollection_url(t *testing.T) {
	check := func(url any, expected bool) {
		isCollection, err := IsCollection(url)
		assert.Nil(t, err)
		assert.Equal(t, isCollection, expected)
	}

	check("https://steamcommunity.com/workshop/filedetails/?id=1884025115", true)
	check("https://steamcommunity.com/workshop/filedetails/?id=818773962", false)
}
func Test_IsCollection_id(t *testing.T) {
	check := func(url any, expected bool) {
		isCollection, err := IsCollection(url)
		assert.Nil(t, err)
		assert.Equal(t, isCollection, expected)
	}

	check(1884025115, true)
	check(818773962, false)
}

func Test_IsCollection_err(t *testing.T) {
	check := func(a any) {
		_, err := IsCollection(a)
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, errors.ErrType))
	}

	check(nil)
	check(1.2)
	check('r')
}
