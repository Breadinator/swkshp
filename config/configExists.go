package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateConfigIfNotExists() [2]error {
	var errs = *new([2]error)
	def := GetConfigDefault()

	errs[0] = makeJSONIfNotExists(def.Paths.Main, def.Main)
	errs[1] = makeJSONIfNotExists(def.Paths.Games, def.Games)

	return errs
}

func makeJSONIfNotExists(path string, jsonStruct any) error {
	os.MkdirAll(filepath.Dir(path), 0770)
	_, err := os.Stat(path)
	if err != nil {
		data, err := json.Marshal(&jsonStruct)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(path, data, 0666)
	}
	return nil
}
