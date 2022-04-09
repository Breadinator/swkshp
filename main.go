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
	if !setup() {
		utils.Info("Setup failed. Exiting.")
		return
	}
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

	checkUpdate()
	return true
}

func close() {
	// close all db connections, log all errors
	utils.Errs(versions.DBCloseAll())
}
