package config

import "github.com/breadinator/swkshp/utils"

func SetVersion(version string) error {
	path, err := GetConfigMain()
	if err != nil {
		return err
	}
	createIfNotExists(path, true)
	return utils.JSONSet(path, "version", version)
}
