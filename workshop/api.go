package workshop

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/parnurzeal/gorequest"
)

const ENDPOINT string = "https://node05.steamworkshopdownloader.io/prod//api/"

// Checks if a resource is available.
//
// Param `resource` is an int for its Workshop ID or a string for its Workshop URL.
func CheckAvailable(resource interface{}) (string, []error) {
	var id int
	switch t := resource.(type) {
	case int:
		id = t
	case string:
		var err error
		id, err = WorkshopIDFromURL(t)
		if err != nil {
			return "", []error{errors.New("couldn't parse resource")}
		}
	default:
		return "", []error{errors.New("couldn't parse resource")}
	}

	downloadFormat := "raw"

	request := gorequest.New()
	resp, body, errs := request.Post(ENDPOINT+"download/request").
		Set("Content-Type", "application/json").
		Send(`{"publishedFileId":` + strconv.Itoa(id) + `, "collectionId":null, "hidden":false, "downloadFormat":"` + downloadFormat + `", "autodownload":true}`).
		End()

	if resp.StatusCode != 200 {
		return "", []error{fmt.Errorf("game not avalable or server is down. Status code: %d", resp.StatusCode)}
	}

	return body, errs
}
