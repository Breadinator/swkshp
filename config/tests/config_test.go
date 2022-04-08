package config_tests

import (
	"fmt"
	"testing"

	"github.com/breadinator/swkshp/config"
	"github.com/stretchr/testify/assert"
)

func Test_GetConfig(t *testing.T) {
	conf, err := config.GetConfig()
	assert.Nil(t, err)
	fmt.Printf("%+v\n", conf)
}

func Test_SetConfig(t *testing.T) {
	conf, err := config.GetConfig()
	assert.Nil(t, err)
	defConf := config.GetConfigDefault()
	conf.Main.FileReadBuffer = defConf.Main.FileReadBuffer
	config.SaveConfig(conf)
}
