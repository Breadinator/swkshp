package main

import "github.com/breadinator/swkshp/config"

func autoUpdate() {
	// TODO check for new version
	// TODO install new version? maybe just check

	// updates the config version to be what is listed in main.go
	// TODO check if version already set, so it doesn't overwrite with the same content
	config.SetVersion(VERSION)
}
