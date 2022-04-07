package workshop

import (
	"errors"
	"fmt"
	"os"

	"github.com/breadinator/swkshp/utils"
	"github.com/breadinator/swkshp/versions"
)

// Extracts things that match "dir/id - (downloaded title).zip" where dir is a given directory and id is the given Workshop ID.
//
// Deletes the archive if successful.
func ExtractResource(id int, dir, game string, verbose bool) (string, error) {
	dbEntry := versions.Entry{
		ID: int64(id),
	}

	url, _ := WorkshopIDToURL(id)
	title, err := GetResourceTitle(url)
	if err != nil {
		return "", err
	}
	if updated, ok := utils.ParseWorkshopTimestamp(url); ok {
		dbEntry.Updated = updated
	}

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

	dbEntry.Sum, _ = utils.GetFileMD5(zipped, 512)

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
