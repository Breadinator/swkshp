package main

import (
	"github.com/breadinator/swkshp/cmd"
	"github.com/breadinator/swkshp/config"
	"github.com/breadinator/swkshp/utils"
	"github.com/breadinator/swkshp/versions"
)

const (
	VERSION = config.VERSION
)

func main() {
	setup()
	defer close()
	cmd.Execute()
}

func setup() bool {
	var err error
	config.Conf, err = config.GetConfig()
	if err != nil {
		utils.Err(err)
		return false
	}

	autoUpdate()

	return true
}

func close() {
	// close all databases
	for _, err := range versions.DBCloseAll() {
		utils.Err(err)
	}
}
