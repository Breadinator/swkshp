package workshop

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const workshop string = `https://steamcommunity.com/workshop/filedetails/?id=%s`

// Converts a Workshop ID to a Workshop URL
func WorkshopIDToURL(id any) (string, bool) {
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

// Regular expression to get the ID from a URL.
var getIDRegexp = regexp.MustCompile(`\?id=(\d+)`) // stricter: `steamcommunity\.com\/(?:workshop|sharedfiles)\/filedetails\/\?id=(\d+)`

// Extracts the Workshop ID from its URL.
func WorkshopIDFromURL(url string) (int, error) {
	matches := getIDRegexp.FindStringSubmatch(url)

	if len(matches) <= 1 {
		return 0, errors.New("no match")
	}

	return strconv.Atoi(matches[1])
}
