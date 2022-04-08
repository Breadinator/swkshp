package downloader

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/breadinator/swkshp/config"
	"github.com/breadinator/swkshp/resource"
	"github.com/breadinator/swkshp/utils"
	"github.com/breadinator/swkshp/versions"
	"github.com/breadinator/swkshp/workshop"
	"github.com/spf13/cobra"
)

// Downloads and extracts from Steam Workshop to the mods folder.
func DefaultExtract(cmd *cobra.Command, args []string) {
	game, err := cmd.Flags().GetString("game")
	if err != nil {
		utils.Err(err)
		return
	}
	if game == "" {
		game, err = workshop.GetGame(args)
		if err != nil {
			utils.Err(err)
			return
		}
	}
	game = strings.ToLower(game)

	modFolder, ok := config.Conf.Games[game] //config.GetGame(game)
	if !ok || modFolder == "" {
		utils.Info("Please set the mod folder for %s using:\n	swkshp.exe config game \"%s\" \"C:/path/to/mod/folder\"", game, game)
		return
	}

	url := strings.Join(args, " ")

	if isCollection, err := workshop.IsCollection(url); isCollection {
		resp, err := http.Get(url)
		if err != nil {
			utils.Err(err)
			return
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			utils.Err(err)
			return
		}

		doc.Find("div.collectionItem").Each(func(i int, s *goquery.Selection) {
			link, exists := s.Find("div.collectionItemDetails a").Attr("href")
			if exists {
				defaultExtractResource(link, modFolder, game)
			}
		})

	} else if err != nil {
		utils.Err(err)
		return
	} else {
		defaultExtractResource(url, modFolder, game)
	}

}

func defaultExtractResource(url, dir, game string) {
	r := resource.ResourceFromURL(url)
	id, err := r.ID()
	if err != nil {
		utils.Err(err)
		return
	}

	// checks if newer version detected
	entry, err := versions.GetModEntry(game, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		utils.Err(err)
		return
	}

	if t, err := r.Updated(); err != nil {
		utils.Err(err)
		return
	} else if t.Before(entry.Updated) {
		return
	}

	if _, err := workshop.ExtractResource(r, dir, game, true); err != nil {
		utils.Err(err)
	}
}
