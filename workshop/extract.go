package workshop

import (
	"fmt"
	"os"

	"github.com/breadinator/swkshp/utils"
)

func ExtractResource(id int, dir string, verbose bool) (string, error) {
	url, _ := WorkshopIDToURL(id)
	title, err := GetResourceTitle(url)
	if err != nil {
		return "", err
	}

	unzipped := fmt.Sprintf("%s%c%d - %s", dir, os.PathSeparator, id, title)
	zipped := unzipped + ".zip"
	f, err := os.OpenFile(zipped, os.O_CREATE, 0666)
	if err != nil {
		if verbose {
			fmt.Println(err)
		}
		return "", err
	}

	if verbose {
		fmt.Printf("Downloading %d...\n", id)
	}
	DownloadResource(id, f)
	f.Close()

	if verbose {
		fmt.Printf("Extracting %s...\n", zipped)
	}
	utils.Unzip(zipped, unzipped)

	if verbose {
		fmt.Printf("Deleting %s...\n", zipped)
	}
	utils.Delete(zipped)

	return unzipped, nil
}
