package workshop

import (
	"github.com/breadinator/swkshp/utils"
)

// Returns true if the given page is a URL.
//
// At the top of a Workshop page theres "breadcrumbs" which look like "Game > Workshop > Collections > User's Workshop".
// This function checks if "Collection" is in those breadcrumbs.
func IsCollection(url string) (bool, error) {
	breadcrumbs, err := GetBreadcrumbs(url)
	if err != nil {
		return false, err
	}
	return utils.In(breadcrumbs, "Collections"), nil
}
