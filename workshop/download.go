package workshop

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/breadinator/swkshp/errors"
	"github.com/parnurzeal/gorequest"
	"github.com/tidwall/gjson"
)

func downloadFile(url string, output io.Writer) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(output, res.Body)
	return err
}

// Writes Workshop resource by ID to given Writer.
func DownloadResource(id int, output io.Writer) error {
	tries := 10

	url := fmt.Sprintf(`https://steamcommunity.com/sharedfiles/filedetails/?id=%d`, id)
	uuidRaw, errs := CheckAvailable(url)

	if errs != nil {
		return errs[0]
	}

	uuid := gjson.Get(uuidRaw, "uuid").String()

	ready := false
	node := ""
	path := ""

	request := gorequest.New()

	for i := 0; i < tries; i++ {
		_, resp, _ := request.Post(ENDPOINT+"download/status").
			Set("Content-Type", "application/json").
			Send(fmt.Sprintf(`{"uuids": ["%s"]}`, uuid)).
			End()

		if strings.Contains(resp, "prepared") {
			ready = true
			node = gjson.Get(resp, uuid+".storageNode").String()
			path = gjson.Get(resp, uuid+".storagePath").String()
			break
		}
	}

	if !ready {
		return errors.ErrResourceNotReady
	}

	url = fmt.Sprintf(`https://%s/prod//storage/%s?uuid=%s`, node, path, uuid)
	return downloadFile(url, output)
}
