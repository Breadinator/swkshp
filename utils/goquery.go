package utils

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Deprecated: use workshop.GetBreadcrumbs
func GetBreadcrumbs(url string) ([]string, error) {
	doc, err := GetDoc(url)
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

// Gets the goquery document of a given web page.
func GetDoc(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return goquery.NewDocumentFromReader(resp.Body)
}

// Gets the Updated timestamp
func GetUpdated(url string) (string, error) {
	doc, err := GetDoc(url)
	if err != nil {
		return "", err
	}
	return doc.Find("div.detailsStatRight").Last().Html()
}
