package resource

import (
	"errors"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/breadinator/swkshp/utils"
	"golang.org/x/exp/constraints"
)

type Resource struct {
	id          int
	url         string
	doc         *goquery.Document
	updated     time.Time
	breadcrumbs []string
	title       string
	timestamp   string
}

func ResourceFromID[T constraints.Integer](id T) Resource {
	return Resource{id: int(id)}
}

func ResourceFromURL[T ~string](url T) Resource {
	return Resource{url: string(url)}
}

func (r *Resource) ID() (int, error) {
	var err error
	if r.id == 0 {
		r.id, err = utils.WorkshopIDFromURL(r.url)
	}
	return r.id, err
}

func (r *Resource) URL() (string, bool) {
	ok := true
	if r.url == "" {
		r.url, ok = utils.WorkshopIDToURL(r.id)
	}
	return r.url, ok
}

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

func (r *Resource) Breadcrumbs() ([]string, error) {
	var err error
	if len(r.breadcrumbs) == 0 {
		url, _ := r.URL()
		r.breadcrumbs, err = utils.GetBreadcrumbs(url)
	}
	return r.breadcrumbs, err
}

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
