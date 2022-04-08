package cmd

import (
	"errors"
	"os"

	"github.com/breadinator/swkshp/config"
	"github.com/breadinator/swkshp/utils"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config game <game> <download_path>",
	Short: "Configure swkshp",
	Long:  `Configure swkshp`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 { // Get config path and open it
			p, e := config.GetConfigPath()
			if e != nil {
				utils.Err(e)
				panic(e)
			}

			utils.Info("Config stored at: " + p)

			if _, e := os.Stat(p); errors.Is(e, os.ErrNotExist) {
				if err := os.Mkdir(p, os.ModePerm); err != nil {
					utils.Err(err)
					panic(err)
				}
			}

			open.Run(p)

		} else { // Set config
			if args[0] == "game" {
				if len(args) == 1 {
					utils.Info("Config for games stored at " + config.Conf.Paths.Games)
				} else if len(args) == 2 { // get game location
					path, ok := config.Conf.Games[args[1]]
					if ok {
						utils.Info("For %s, downloading to %s", args[1], path)
					} else {
						utils.Info("Install directory not set for %s", args[1])
					}

				} else if len(args) == 3 { // set game location
					config.Conf.Games[args[1]] = args[2]
					if err := config.SaveConfig(config.Conf); err != nil {
						utils.Err(err)
						panic(err)
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
