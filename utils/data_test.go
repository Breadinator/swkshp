package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WorkshopIDToURLInt(t *testing.T) {
	url, ok := WorkshopIDToURL(123)
	assert.True(t, ok)
	assert.Equal(t, "steamcommunity.com/sharedfiles/filedetails/?id=123", url)
}

func Test_WorkshopIDToURLStr(t *testing.T) {
	url, ok := WorkshopIDToURL("123")
	assert.True(t, ok)
	assert.Equal(t, "steamcommunity.com/sharedfiles/filedetails/?id=123", url)
}

func Test_WorkshopIDToURLFails(t *testing.T) {
	var ok bool

	// Bool
	_, ok = WorkshopIDToURL(false)
	assert.False(t, ok)

	// Nil pointer
	_, ok = WorkshopIDToURL(nil)
	assert.False(t, ok)

	// Floating point
	_, ok = WorkshopIDToURL(7.23)
	assert.False(t, ok)

	// Array
	_, ok = WorkshopIDToURL([...]int{0, 1})
	assert.False(t, ok)
}

func Test_WorkshopIDFromURLWorking(t *testing.T) {
	id, err := WorkshopIDFromURL("steamcommunity.com/sharedfiles/filedetails/?id=123")
	assert.Nil(t, err)
	assert.Equal(t, 123, id)
}

func Test_WorkshopIDFromURLFail(t *testing.T) {
	_, err := WorkshopIDFromURL("steamcommunity.com/sharedfiles/filedetails/?id= invalid")
	assert.Error(t, err)
}

func Test_Sanitise(t *testing.T) {
	var tests = [...][2]string{
		{"all_valid_characters", "all_valid_characters"},
		{"résumé", "resume"},
		{"áèîçÍÀÔÇ", "aeiciaoc"},
		{"#<&+%>!`&*'|{?\"=}/:\\@", ""},
		{`éVÈRŷ%!\th iNg <t>ogETh&èr`, "everyth_ing_together"},
	}

	for _, test := range tests {
		s, err := Sanitise(test[0])
		assert.Nil(t, err)
		assert.Equal(t, test[1], s)
	}
}
