package workshop

import (
	"github.com/breadinator/swkshp/utils"
	"golang.org/x/exp/constraints"
)

// Converts a Workshop ID to a Workshop URL
func WorkshopIDToURL[T constraints.Integer](id T) (string, bool) {
	return utils.WorkshopIDToURL(int(id))
}

// Extracts the Workshop ID from its URL.
func WorkshopIDFromURL[T ~string](url T) (int, error) {
	return utils.WorkshopIDFromURL(string(url))
}
