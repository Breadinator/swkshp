package resource

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ResourceFromID(t *testing.T) {
	r := ResourceFromID(818773962)
	checkHugsLib(t, &r)
}
func Test_ResourceFromURL(t *testing.T) {
	r := ResourceFromURL("https://steamcommunity.com/workshop/filedetails/?id=818773962")
	checkHugsLib(t, &r)
}

func checkHugsLib(t *testing.T, r *Resource) {
	id, err := r.ID()
	assert.Nil(t, err)
	assert.Equal(t, 818773962, id)

	url, ok := r.URL()
	assert.True(t, ok)
	assert.Equal(t, "https://steamcommunity.com/workshop/filedetails/?id=818773962", url)

	doc, err := r.Doc()
	assert.NotNil(t, doc)
	assert.Nil(t, err)
	if doc != nil {
		html, err := doc.Html()
		assert.Nil(t, err)
		assert.NotEqual(t, "", html)
	}

	updated, err := r.Updated()
	assert.Nil(t, err)
	assert.Equal(t, 2021, updated.Year())

	breadcrumbs, err := r.Breadcrumbs()
	assert.Nil(t, err)
	assert.True(t, len(breadcrumbs) != 0)
	if len(breadcrumbs) == 3 {
		assert.Equal(t, "RimWorld", breadcrumbs[0])
		assert.Equal(t, "Workshop", breadcrumbs[1])
		assert.Equal(t, "UnlimitedHugs&#39;s Workshop", breadcrumbs[2])
	}
}
