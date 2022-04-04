package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/breadinator/swkshp/config"
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
				panic(e)
			}

			fmt.Println("Config stored at: " + p)

			if _, e := os.Stat(p); errors.Is(e, os.ErrNotExist) {
				if err := os.Mkdir(p, os.ModePerm); err != nil {
					panic(err)
				}
			}

			open.Run(p)

		} else { // Set config
			if args[0] == "game" {
				if len(args) == 1 {
					if p, e := config.GetConfigPathGame(); e == nil {
						fmt.Println("Config for games stored at " + p)
					}

				} else if len(args) == 2 { // get game location
					path, ok := config.GetGame(args[1])
					if ok {
						fmt.Printf("For %s, downloading to %s\n", args[1], path)
					} else {
						fmt.Printf("Install directory not set for %s\n", args[1])
					}

				} else if len(args) == 3 { // set game location
					if err := config.SetGame(args[1], args[2]); err != nil {
						panic(err)
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
