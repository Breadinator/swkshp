package downloader

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/breadinator/swkshp/config"
	"github.com/breadinator/swkshp/workshop"
	"github.com/spf13/cobra"
)

// Downloads and extracts from Steam Workshop to the mods folder.
func DefaultExtract(cmd *cobra.Command, args []string) {
	game, err := cmd.Flags().GetString("game")
	if err != nil {
		panic(err)
	}
	if game == "" {
		game, err = workshop.GetGame(args)
		if err != nil {
			panic(err)
		}
	}
	game = strings.ToLower(game)

	modFolder, ok := config.GetGame(game)
	if !ok || modFolder == "" {
		fmt.Printf("Please set the mod folder for %s using:\n	swkshp.exe config game \"%s\" \"C:/path/to/mod/folder\"\n", game, game)
		os.Exit(0)
	}

	url := strings.Join(args, " ")

	if isCollection, err := workshop.IsCollection(url); isCollection {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			panic(err)
		}

		doc.Find("div.collectionItem").Each(func(i int, s *goquery.Selection) {
			link, exists := s.Find("div.collectionItemDetails a").Attr("href")
			if exists {
				defaultExtractResource(link, modFolder)
			}
		})

	} else if err != nil {
		panic(err)
	} else {
		defaultExtractResource(url, modFolder)
	}

}

func defaultExtractResource(url, dir string) {
	id, err := workshop.WorkshopIDFromURL(url)
	if err != nil {
		panic(err)
	}
	if _, err := workshop.ExtractResource(id, dir, true); err != nil {
		panic(err)
	}
}
