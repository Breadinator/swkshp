package resource

import (
	"errors"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/breadinator/swkshp/utils"
	"golang.org/x/exp/constraints"
)

// Represents all the data to be fetched related to a given mod.
// Used to cache the data, instead of repeatedly accessing it again and again.
//
// Initialise using ResourceFromID or ResourceFromURL.
//
// Get values using the methods named after the fields.
type Resource struct {
	id          int               // Mod ID.
	url         string            // URL to the workshop mod page.
	doc         *goquery.Document // Cached goquery web page.
	timestamp   string            // When last updated, according to the webpage, before any parsing.
	updated     time.Time         // When last updated, according to the web page.
	breadcrumbs []string          // The stuff above the workshop title, e.g. "Game> Workshop > User's Workshop" for a single entry.
	title       string            // The name of the mod.
}

func ResourceFromID[T constraints.Integer](id T) Resource {
	return Resource{id: int(id)}
}

func ResourceFromURL[T ~string](url T) Resource {
	return Resource{url: string(url)}
}

// If and ID has been given, it will return that.
// Otherwise, it will attempt to parse it from the given URL.
func (r *Resource) ID() (int, error) {
	var err error
	if r.id == 0 {
		r.id, err = utils.WorkshopIDFromURL(r.url)
	}
	return r.id, err
}

// If a URL has been given, it will return that.
// Otherwise, it will construct one using the given ID.
func (r *Resource) URL() (string, bool) {
	ok := true
	if r.url == "" {
		r.url, ok = utils.WorkshopIDToURL(r.id)
	}
	return r.url, ok
}

// If already fetched, it'll reuse the version stored in memory.
// Otherwise, it'll download the page, parse it into a Goquery document and store that in memory, then give it to you.
func (r *Resource) Doc() (*goquery.Document, error) {
	var err error
	if r.doc == nil {
		url, ok := r.URL()
		if ok {
			r.doc, err = utils.GetDoc(url)
		} else {
			err = errors.New("couldn't get resource URL")
		}
	}
	return r.doc, err
}

// Gets the timestamp of when last updated as a string, unparsed.
func (r *Resource) Timestamp() (string, error) {
	var err error

	if r.timestamp == "" {
		var doc *goquery.Document
		doc, err = r.Doc()
		if doc != nil && err == nil {
			r.timestamp, err = doc.Find("div.detailsStatRight").Last().Html()
		}
	}

	return r.timestamp, err
}

// Parses the string timestamp gotten from Timestamp()
func (r *Resource) Updated() (time.Time, error) {
	var err error
	var nilTime time.Time

	if r.updated == nilTime {
		var timestamp string
		timestamp, err = r.Timestamp()
		if timestamp != "" && err == nil {
			ok := false
			r.updated, ok = utils.ParseWorkshopTimestamp(timestamp)
			if !ok {
				err = fmt.Errorf("couldn't parse timestamp: %s", timestamp)
			}
		}
	}

	return r.updated, err
}

// The stuff above the workshop title, e.g. "Game> Workshop > User's Workshop" for a single entry.
func (r *Resource) Breadcrumbs() ([]string, error) {
	var err error
	if len(r.breadcrumbs) == 0 {
		url, _ := r.URL()
		r.breadcrumbs, err = utils.GetBreadcrumbs(url)
	}
	return r.breadcrumbs, err
}

// Gets the name of the mod from the web page's title header.
func (r *Resource) Title() (string, error) {
	var err error

	if r.title == "" {
		var doc *goquery.Document
		doc, err = r.Doc()

		if err == nil {
			r.title, err = doc.Find(`div.workshopItemTitle`).Html()
		}
	}

	return r.title, err
}
