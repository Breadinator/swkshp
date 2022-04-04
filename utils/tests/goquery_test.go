package utils_tests

import (
	"testing"

	"github.com/breadinator/swkshp/utils"
	"github.com/stretchr/testify/assert"
)

func Test_GetBreadcrumbs(t *testing.T) {
	breadcrumbs, err := utils.GetBreadcrumbs(`https://steamcommunity.com/workshop/filedetails/?id=1884025115`)
	assert.Nil(t, err)
	assert.Len(t, breadcrumbs, 4)

	expected := []string{"RimWorld", "Workshop", "Collections", "Oskar Potocki&#39;s Workshop"}
	for i := 0; i < len(breadcrumbs); i++ {
		assert.Equal(t, expected[i], breadcrumbs[i])
	}
}
