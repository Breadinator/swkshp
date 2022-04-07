package cmd

import (
	"os"

	"github.com/breadinator/swkshp/downloader"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "swkshp [url]",
	Short: "Download from Steam Workshop without authentication",
	Long:  `Download from Steam Workshop without authentication.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			downloader.Download(cmd, args)
		}
	},
	Args: cobra.MaximumNArgs(1),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("game", "g", "", "Specify a game to download for.")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
