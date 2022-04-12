package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/breadinator/swkshp/errors"
	"golang.org/x/exp/constraints"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

/* STRING HANDLING */

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

var (
	illegalChars = regexp.MustCompile(`[^a-zA-Z0-9_]`) // Anything other than either case a-z, arabic numerals or underscores
	wspace       = regexp.MustCompile(`\s`)            // Any whitespace
)

// Replaces accents (e.g. 'é' -> 'e'). Converts to lowercase. Replaces whitespace with underscores. Removes all illegal characters.
func Sanitise[S ~string](str S) (string, error) {
	// Transform accents
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, err := transform.String(t, string(str))

	// Lower case
	s = strings.ToLower(s) // can maybe merge to above

	// Whitespace
	s = wspace.ReplaceAllString(s, "_")

	// Remove illegal characters
	s = illegalChars.ReplaceAllString(s, "")
	return s, err
}

// Regular expression to get the ID from a URL.
var getIDRegexp = regexp.MustCompile(`\?id=(\d+)`) // stricter: `steamcommunity\.com\/(?:workshop|sharedfiles)\/filedetails\/\?id=(\d+)`

// Extracts the Workshop ID from its URL.
func WorkshopIDFromURL(url string) (int, error) {
	matches := getIDRegexp.FindStringSubmatch(url)

	if len(matches) <= 1 {
		return 0, errors.ErrParsingFailed
	}

	return strconv.Atoi(matches[1])
}

/* SLICES DATA HANDLING */

// Checks if the items and order of two slices are equal
func SlicesEqual[T constraints.Ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, item := range a {
		if b[i] != item {
			return false
		}
	}
	return true
}

// Checks if a slice of strings contains an given string.
func In[T comparable](slice []T, item T) bool {
	for _, a := range slice {
		if item == a {
			return true
		}
	}
	return false
}

/* INT HANDLING */

// Pass in any non-zero number of orderables and get back the one that's "smallest" (i.e. it < the rest).
//
// Uses args `a` and `numbers` to ensure at least one value is given.
func Min[T constraints.Ordered](a T, numbers ...T) T {
	smallest := a

	for _, n := range numbers {
		if n < smallest {
			smallest = n
		}
	}

	return smallest
}
