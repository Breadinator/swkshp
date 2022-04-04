package workshop

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/PuerkitoBio/goquery"
)

// Gets the game from a steam workshop ID or URL. Arg "id" accepts int or string, where the former is the ID and the latter is the full URL.
func GetGame(id interface{}) (string, error) {
	switch t := id.(type) {
	case int:
		return getGameInt(t)
	case string:
		return getGameStr(t)
	default:
		return "", fmt.Errorf("expected string or int, got %s", reflect.TypeOf(id))
	}
}

func getGameInt(id int) (string, error) {
	a, _ := WorkshopIDToURL(id)
	return getGameStr(a)
}

func getGameStr(url string) (string, error) {
	// TODO: check if valid URL

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	return doc.Find(`div.breadcrumbs a`).First().Html()
}
