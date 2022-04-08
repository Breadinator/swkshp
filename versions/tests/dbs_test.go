package versions_tests

import (
	"testing"

	"github.com/breadinator/swkshp/versions"
	"github.com/stretchr/testify/assert"
)

func Test_DBOpen_DBCloseAll(t *testing.T) {
	assert.True(t, versions.DBLen() == 0)

	db, err := versions.DBOpen("rimworld")
	assert.NotNil(t, db)
	assert.Nil(t, err)

	errs := versions.DBCloseAll()
	assert.True(t, len(errs) == 0, "%v should be empty", errs)
	assert.True(t, versions.DBLen() == 0)
}
