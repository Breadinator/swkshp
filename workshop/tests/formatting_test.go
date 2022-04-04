package workshop_tests

import (
	"testing"

	"github.com/breadinator/swkshp/workshop"
	"github.com/stretchr/testify/assert"
)

func Test_WorkshopIDToURLInt(t *testing.T) {
	url, ok := workshop.WorkshopIDToURL(123)
	assert.True(t, ok)
	assert.Equal(t, "steamcommunity.com/sharedfiles/filedetails/?id=123", url)
}

func Test_WorkshopIDToURLStr(t *testing.T) {
	url, ok := workshop.WorkshopIDToURL("123")
	assert.True(t, ok)
	assert.Equal(t, "steamcommunity.com/sharedfiles/filedetails/?id=123", url)
}

func Test_WorkshopIDToURLFails(t *testing.T) {
	var ok bool

	// Bool
	_, ok = workshop.WorkshopIDToURL(false)
	assert.False(t, ok)

	// Nil pointer
	_, ok = workshop.WorkshopIDToURL(nil)
	assert.False(t, ok)

	// Floating point
	_, ok = workshop.WorkshopIDToURL(7.23)
	assert.False(t, ok)

	// Array
	_, ok = workshop.WorkshopIDToURL([...]int{0, 1})
	assert.False(t, ok)
}

func Test_WorkshopIDFromURLWorking(t *testing.T) {
	id, err := workshop.WorkshopIDFromURL("steamcommunity.com/sharedfiles/filedetails/?id=123")
	assert.Nil(t, err)
	assert.Equal(t, 123, id)
}

func Test_WorkshopIDFromURLFail(t *testing.T) {
	_, err := workshop.WorkshopIDFromURL("steamcommunity.com/sharedfiles/filedetails/?id= invalid")
	assert.Error(t, err)
}
