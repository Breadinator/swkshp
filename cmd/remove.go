package cmd

import (
	"os"
	"strconv"

	"github.com/breadinator/swkshp/errors"
	"github.com/breadinator/swkshp/resource"
	"github.com/breadinator/swkshp/utils"
	"github.com/breadinator/swkshp/versions"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove <url>",
	Short: "Uninstalls a mod.",
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		game, _ := cmd.Flags().GetString("game")

		// checks if the given
		id, err := strconv.Atoi(url)
		if err != nil {
			id, err = utils.WorkshopIDFromURL(url)
			if err != nil {
				utils.Err(err)
				return
			}
		}

		r := resource.ResourceFromURL(url)

		// gets game if none provided
		if game == "" {
			game, err = r.Game()
			if err != nil {
				utils.Err(err)
				return
			} else if game == "" {
				utils.Err(errors.Wrap(errors.ErrGameUnavailable, "game with ID %d unavailable", id))
				utils.Info("You can provide your own game info using the -g flag")
				return
			}
		}

		utils.Info("Removing %d...", id)

		// checks the database for the entry
		ent, err := versions.GetModEntry(game, id)
		if err != nil {
			utils.Err(err, "Error getting database entry for", strconv.Itoa(id), "in", game)
			return
		}

		// if the directory still exists, remove it
		if _, err := os.Stat(ent.Path); !errors.Is(err, os.ErrNotExist) {
			err = os.RemoveAll(ent.Path)
			if err != nil {
				utils.Err(err)
				return
			}
			utils.Info("Deleted mod files")
		}

		// remove entry from database
		if _, err = versions.RemoveModEntry(game, id); err != nil {
			utils.Err(err)
			return
		}
		utils.Info("Mod removed from internal database")

		utils.Info("Successfully removed %d.", id)
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.Flags().StringP("game", "g", "", "Specify a game to download for.")
}
