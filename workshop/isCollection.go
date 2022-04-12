package workshop

import (
	"reflect"

	"github.com/breadinator/swkshp/errors"
	"github.com/breadinator/swkshp/resource"
	"github.com/breadinator/swkshp/utils"
)

// Returns true if the given page is a URL.
//
// At the top of a Workshop page theres "breadcrumbs" which look like "Game > Workshop > Collections > User's Workshop".
// This function checks if "Collection" is in those breadcrumbs.
func IsCollection(resrc any) (bool, error) {
	var r resource.Resource
	switch t := resrc.(type) {
	case int:
		r = resource.ResourceFromID(t)
	case string:
		r = resource.ResourceFromURL(t)
	case resource.Resource:
		r = t
	default:
		return false, errors.Wrap(errors.ErrType, "expected int, string or Resource, got %s", reflect.TypeOf(resrc))
	}

	breadcrumbs, err := r.Breadcrumbs()
	if err != nil {
		return false, err
	}
	return utils.In(breadcrumbs, "Collections"), nil
}
