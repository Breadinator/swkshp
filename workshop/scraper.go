package workshop

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/breadinator/swkshp/utils"
)

// Gets the Workshop title from the page given.
func GetResourceTitle(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", errors.New(res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	return doc.Find("div.workshopItemTitle").Text(), nil
}

// At the top of a Workshop page theres "breadcrumbs" which look like "Game > Workshop > Collections > User's Workshop".
// This function looks for thos "breadcrumbs" then gives them like []string{ "Game", "Workshop", "Collections", "User's Workshop" }
func GetBreadcrumbs(url string) ([]string, error) {
	doc, err := utils.GetDoc(url)
	if err != nil {
		return nil, err
	}

	breadcrumbs := make([]string, 0)

	var eachErr error
	doc.Find(`div.breadcrumbs a`).Each(func(_ int, s *goquery.Selection) {
		breadcrumb, err := s.Html()
		if err != nil {
			eachErr = err
			return
		}
		breadcrumbs = append(breadcrumbs, breadcrumb)
	})

	return breadcrumbs, eachErr
}
