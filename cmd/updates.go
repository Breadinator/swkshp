package cmd

import (
	"fmt"
	"strconv"

	"github.com/breadinator/swkshp/resource"
	"github.com/breadinator/swkshp/utils"
	"github.com/breadinator/swkshp/versions"
	"github.com/spf13/cobra"
)

// updatesCmd represents the updates command
var updatesCmd = &cobra.Command{
	Use:   "updates",
	Short: "Check if any updates are available for a given game, or all games if no game provided.",
	Long:  `Check if any updates are available for a given game, or all games if no game provided.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		game, err := cmd.Flags().GetString("game")
		if err != nil {
			utils.Err(err)
			return
		}

		games := make(map[string]map[*resource.Resource]bool)
		var errs []error

		if game == "" {
			games, errs = versions.CheckForUpdate()
		} else {
			games, errs = versions.CheckForUpdate(game)
		}

		utils.Errs(errs)

		if len(games) == 0 {
			utils.Info("No updates available.")
			return
		}

		// i.e. looking for a single game
		if game != "" {
			resources, ok := games[game]
			if !ok || len(resources) == 0 {
				utils.Info("No updates available.")
				return
			}
			updates := whatUpdates(resources)
			if len(updates) == 0 {
				utils.Info("No updates available.")
			} else {
				msg := fmt.Sprintf("Updates for %s:", game)
				for _, resource := range updates {
					title, err := resource.Title()
					if err != nil {
						utils.Err(err)
						continue
					}
					id, err := resource.ID()
					if err != nil {
						utils.Err(err)
						continue
					}

					msg += "\n         * " + title + " (" + strconv.Itoa(id) + ")"
				}
				utils.Info(msg)
			}
			return
		}

		// checking all games
		gamesWithUpdates := make([]string, 0)
		for game, resources := range games {
			for _, updateAvailable := range resources {
				if updateAvailable {
					gamesWithUpdates = append(gamesWithUpdates, game)
					break
				}
			}
		}

		if len(gamesWithUpdates) == 0 {
			utils.Info("No updates available.")
			return
		}

		msg := "Updates available for the following games:"
		for _, game := range gamesWithUpdates {
			msg += "\n         * " + game
		}
		utils.Info(msg)
	},
}

func init() {
	rootCmd.AddCommand(updatesCmd)
	updatesCmd.Flags().StringP("game", "g", "", "Specify a game to download for.")
}

func whatUpdates(resources map[*resource.Resource]bool) []*resource.Resource {
	resourcesWithUpdates := make([]*resource.Resource, 0)

	for resource, updatesAvailable := range resources {
		if updatesAvailable {
			resourcesWithUpdates = append(resourcesWithUpdates, resource)
		}
	}

	return resourcesWithUpdates
}
