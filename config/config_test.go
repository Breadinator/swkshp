package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetConfig(t *testing.T) {
	conf, err := GetConfig()
	assert.Nil(t, err)
	fmt.Printf("%+v\n", conf)
}

func Test_SetConfig(t *testing.T) {
	conf, err := GetConfig()
	assert.Nil(t, err)
	defConf := GetConfigDefault()
	conf.Main.FileReadBuffer = defConf.Main.FileReadBuffer
	SaveConfig(conf)
}
