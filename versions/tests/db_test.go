package versions_tests

import (
	"crypto/md5"
	"fmt"
	"testing"
	"time"

	"github.com/breadinator/swkshp/versions"
	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	fmt.Println(versions.GetDBPath("test", true))
}

var s [16]byte = md5.Sum([]byte("test"))

func TestB(t *testing.T) {
	entry := versions.Entry{
		ID:      123,
		Path:    "C:/example/path",
		Sum:     s[:],
		Updated: time.Now(),
	}
	fmt.Println(versions.UpdateModEntry("test", entry))
}

func TestC(t *testing.T) {
	TestB(t)

	ent, err := versions.GetModEntry("test", 123)
	fmt.Println(*ent, err)

	assert.EqualValues(t, 123, ent.ID)
	assert.Equal(t, "C:/example/path", ent.Path)
	assert.Equal(t, s[:], ent.Sum)
	assert.True(t, ent.Updated.Before(time.Now()))
}

func TestD(t *testing.T) {
	_, err := versions.GetModEntry("test", -1)
	assert.NotNil(t, err)
}

func Test_GetAllEntries(t *testing.T) {
	ents, err := versions.GetAllEntries("test")
	assert.Nil(t, err)
	fmt.Printf("%+v\n", ents)
}
