package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func SetGame(game, directory string) error {
	directory, err := filepath.Abs(directory)
	if err != nil {
		return err
	}

	p, err := GetConfigPathGame()
	if err != nil {
		return err
	}
	createIfNotExists(p)

	bytes, err := ioutil.ReadFile(p)
	if err != nil {
		return err
	}

	games := string(bytes)
	if len(games) == 0 {
		games = "{}"
	}

	games, err = sjson.Set(string(bytes), strings.ToLower(game), directory)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(p, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte(games))

	return nil
}

func GetGame(game string) (string, bool) {
	p, e := GetConfigPathGame()
	if e != nil {
		return "", false
	}
	createIfNotExists(p)

	bytes, err := ioutil.ReadFile(p)
	if err != nil {
		return "", false
	}

	return gjson.Get(string(bytes), strings.ToLower(game)).String(), true
}
