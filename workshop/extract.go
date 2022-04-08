package workshop

import (
	"errors"
	"fmt"
	"os"

	"github.com/breadinator/swkshp/config"
	"github.com/breadinator/swkshp/resource"
	"github.com/breadinator/swkshp/utils"
	"github.com/breadinator/swkshp/versions"
)

// Extracts things that match "dir/id - (downloaded title).zip" where dir is a given directory and id is the given Workshop ID.
//
// Deletes the archive if successful.
func ExtractResource(r resource.Resource, dir, game string, verbose bool) (string, error) {
	id, err := r.ID()
	if err != nil {
		return "", err
	}

	dbEntry := versions.Entry{
		ID: int64(id),
	}

	url, _ := r.URL()
	title, err := GetResourceTitle(url)
	if err != nil {
		return "", err
	}
	dbEntry.Updated, _ = r.Updated()

	unzipped := fmt.Sprintf("%s%c%d - %s", dir, os.PathSeparator, id, title)
	dbEntry.Path = unzipped
	zipped := unzipped + ".zip"
	f, err := os.OpenFile(zipped, os.O_CREATE, 0666)
	if err != nil {
		if verbose {
			utils.Err(err)
		}
		return "", err
	}
	if verbose {
		utils.Info("Downloading %d...", id)
	}
	DownloadResource(id, f)
	f.Close()

	bufSize := config.Conf.Main.FileReadBuffer
	if bufSize == 0 {
		bufSize = config.DefaultConfig.Main.FileReadBuffer
		utils.Warn("Failed to get buffer size. Using default %d.", bufSize)
	}
	dbEntry.Sum, _ = utils.GetFileMD5(zipped, bufSize)

	ent, _ := versions.GetModEntry(game, id)
	_, err = os.Stat(ent.Path)
	if !utils.SlicesEqual(ent.Sum, dbEntry.Sum) || errors.Is(err, os.ErrNotExist) {
		if verbose {
			utils.Info("Extracting %s...", zipped)
		}
		utils.Unzip(zipped, unzipped)
		versions.UpdateModEntry(game, dbEntry)
	} else if verbose {
		utils.Info("Same already exists")
	}

	if verbose {
		utils.Info("Deleting %s", zipped)
	}
	err = utils.Delete(zipped)

	return unzipped, err
}
