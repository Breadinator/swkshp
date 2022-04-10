package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetBreadcrumbs(t *testing.T) {
	breadcrumbs, err := GetBreadcrumbs(`https://steamcommunity.com/workshop/filedetails/?id=1884025115`)
	assert.Nil(t, err)
	assert.Len(t, breadcrumbs, 4)

	expected := []string{"RimWorld", "Workshop", "Collections", "Oskar Potocki&#39;s Workshop"}
	for i := 0; i < len(breadcrumbs); i++ {
		assert.Equal(t, expected[i], breadcrumbs[i])
	}
}

func Test_GetUpdated(t *testing.T) {
	timestamp, err := GetUpdated("https://steamcommunity.com/workshop/filedetails/?id=818773962")
	assert.Nil(t, err)
	assert.Equal(t, "25 Sep, 2021 @ 3:17am", timestamp) //utc time i think?
	fmt.Println(timestamp)

	time, ok := ParseWorkshopTimestamp(timestamp)
	assert.True(t, ok)
	fmt.Println(time)
}
