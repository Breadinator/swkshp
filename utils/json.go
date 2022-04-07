package utils

import (
	"io/ioutil"
	"os"

	"github.com/tidwall/sjson"
)

func JSONSet(file, key string, value any) error {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	new, err := sjson.SetBytes(contents, key, value)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(file, os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(new)
	return nil
}
