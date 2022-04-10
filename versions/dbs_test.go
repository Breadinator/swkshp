package versions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DBOpen_DBCloseAll(t *testing.T) {
	assert.True(t, DBLen() == 0)

	db, err := DBOpen("rimworld")
	assert.NotNil(t, db)
	assert.Nil(t, err)

	errs := DBCloseAll()
	assert.True(t, len(errs) == 0, "%v should be empty", errs)
	assert.True(t, DBLen() == 0)
}
