package workshop

import (
	"fmt"
	"strconv"

	"github.com/breadinator/swkshp/errors"
	"github.com/breadinator/swkshp/utils"
	"github.com/parnurzeal/gorequest"
)

const ENDPOINT string = "https://node05.steamworkshopdownloader.io/prod//api/"

// Checks if a resource is available.
//
// Param `resource` is an int for its Workshop ID or a string for its Workshop URL.
func CheckAvailable(resource any) (string, []error) {
	var id int
	switch t := resource.(type) {
	case int:
		id = t
	case string:
		var err error
		id, err = utils.WorkshopIDFromURL(t)
		if err != nil {
			return "", []error{errors.ErrParsingFailed}
		}
	default:
		return "", []error{errors.ErrParsingFailed}
	}

	downloadFormat := "raw"

	request := gorequest.New()
	resp, body, errs := request.Post(ENDPOINT+"download/request").
		Set("Content-Type", "application/json").
		Send(`{"publishedFileId":` + strconv.Itoa(id) + `, "collectionId":null, "hidden":false, "downloadFormat":"` + downloadFormat + `", "autodownload":true}`).
		End()

	if resp.StatusCode != 200 {
		return "", []error{errors.Wrap(errors.ErrGameNotFound, fmt.Sprintf("Status: %d", resp.StatusCode))}
	}

	return body, errs
}
