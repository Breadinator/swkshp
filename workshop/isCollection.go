package workshop

import (
	"github.com/breadinator/swkshp/utils"
)

func IsCollection(url string) (bool, error) {
	breadcrumbs, err := utils.GetBreadcrumbs(url)
	if err != nil {
		return false, err
	}
	return utils.In(breadcrumbs, "Collections"), nil
}
