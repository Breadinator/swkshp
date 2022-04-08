package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/breadinator/swkshp/utils"
)

func GetConfig() (*Config, error) {
	// var init
	conf := Config{
		Games: make(map[string]string),
	}

	// Load main config
	mainPath, err := GetConfigMain()
	if err != nil {
		return &conf, err
	}
	conf.Paths.Main = mainPath
	mainConf, err := ioutil.ReadFile(mainPath)
	if err != nil {
		return &conf, err
	}
	if err := json.Unmarshal(mainConf, &conf.Main); err != nil {
		return &conf, err
	}

	// Load mod folder directories
	gamesPath, err := GetConfigPathGame()
	if err != nil {
		return &conf, err
	}
	conf.Paths.Games = gamesPath
	gamesConfig, err := ioutil.ReadFile(gamesPath)
	if err != nil {
		return &conf, err
	}

	return &conf, json.Unmarshal(gamesConfig, &conf.Games)
}

func SaveConfig(c *Config) error {
	// var init
	var err error

	// gets config paths if not set
	if utils.IsNilValue(c.Paths.Main) {
		c.Paths.Main, err = GetConfigPathMain()
		if err != nil {
			return err
		}
	}
	if utils.IsNilValue(c.Paths.Games) {
		c.Paths.Games, err = GetConfigPathGame()
		if err != nil {
			return err
		}
	}

	// write to main config file
	mainFile, err := os.OpenFile(c.Paths.Main, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer mainFile.Close()
	mainMarsh, err := json.MarshalIndent(c.Main, "", "    ")
	if err != nil {
		return err
	}
	if _, err := mainFile.Write(mainMarsh); err != nil {
		return err
	}

	// write to games config file
	gamesFile, err := os.OpenFile(c.Paths.Games, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer gamesFile.Close()
	gamesMarsh, err := json.MarshalIndent(c.Games, "", "    ")
	if err != nil {
		return err
	}
	if _, err := gamesFile.Write(gamesMarsh); err != nil {
		return err
	}

	return nil
}
