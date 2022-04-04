package cmd

import (
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download <url>",
	Short: rootCmd.Short,
	Long:  rootCmd.Long,
	Run:   rootCmd.Run,
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringP("game", "g", "", "Specify a game to download for.")
}
