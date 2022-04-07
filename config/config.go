package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func GetConfigPath() (string, error) {
	uconf, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%cswkshp", uconf, os.PathSeparator), nil
}

func GetConfigMain() (string, error) {
	confDir, err := GetConfigPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(confDir, "swkshp.json"), nil
}

func GetConfigPathGame() (string, error) {
	conf, err := GetConfigPath()
	if err != nil {
		return "", err
	}
	return conf + string(os.PathSeparator) + "games.json", nil
}

func createIfNotExists(path string, writeEmptyJSON ...bool) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		_, err = os.Create(path)
		if len(writeEmptyJSON) != 0 && writeEmptyJSON[0] && err != nil {
			f, err := os.Open(path)
			if err != nil {
				return false
			}
			f.WriteString("{}")
			f.Close()
		}
		return err != nil
	}
	return false
}
