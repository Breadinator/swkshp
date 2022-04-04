package config

import (
	"errors"
	"fmt"
	"os"
)

func GetConfigPath() (string, error) {
	uconf, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%cswkshp", uconf, os.PathSeparator), nil
}

func GetConfigPathGame() (string, error) {
	conf, err := GetConfigPath()
	if err != nil {
		return "", err
	}
	return conf + string(os.PathSeparator) + "games.json", nil
}

func createIfNotExists(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		os.Create(path)
	}
}
