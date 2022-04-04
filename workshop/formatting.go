package workshop

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const workshop string = `https://steamcommunity.com/workshop/filedetails/?id=%s`

func WorkshopIDToURL(id interface{}) (string, bool) {
	var idStr string
	switch t := id.(type) {
	case int:
		idStr = strconv.Itoa(t)
	case string:
		idStr = t
	default:
		return "", false
	}

	return fmt.Sprintf(workshop, idStr), true
}

var getIDRegexp = regexp.MustCompile(`\?id=(\d+)`) // stricter: `steamcommunity\.com\/(?:workshop|sharedfiles)\/filedetails\/\?id=(\d+)`

func WorkshopIDFromURL(url string) (int, error) {
	matches := getIDRegexp.FindStringSubmatch(url)

	if len(matches) <= 1 {
		return 0, errors.New("no match")
	}

	return strconv.Atoi(matches[1])
}
