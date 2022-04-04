package downloader

import (
	"github.com/spf13/cobra"
)

func Root(cmd *cobra.Command, args []string) {
	defaultExtract(cmd, args)

	//for special game support, right now not needed
	/*game, err := cmd.Flags().GetString("game")
	if err != nil {
		panic(err)
	}
	if game == "" {
		game, err = workshop.GetGame(strings.Join(args[0:], " "))
		if err != nil {
			panic(err)
		}
	}

	switch strings.ToLower(game) {
	case "rimworld":
		rimworld(cmd, args)
	default:
		fmt.Printf("Game %s doesn't have special support. Using default extract download.\n", game)
		defaultExtract(cmd, args)
	}*/
}
