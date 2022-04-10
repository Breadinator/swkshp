package config

import (
	"os"
	"path/filepath"
)

const VERSION = "v1.3.0"

type Config struct {
	Paths struct{ Main, Games string } // The paths to the config files
	Main  main
	Games map[string]string // The paths to the games
}

type main struct { // Main config file
	Version        string `json:"version"` // SWkshp version
	FileReadBuffer int    `json:"buf"`     // How many bytes in buffers for reading files
}

var Conf, _ = GetConfig()

func GetConfigPath() (string, error) {
	uconf, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(uconf, "swkshp"), nil
}

func GetConfigPathMain() (string, error) {
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
	return filepath.Join(conf, "games.json"), nil
}
