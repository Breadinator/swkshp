package workshop

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

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
