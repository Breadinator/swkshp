/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/breadinator/swkshp/downloader"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download <url>",
	Short: rootCmd.Short,
	Long:  rootCmd.Long,
	Run: func(cmd *cobra.Command, args []string) {
		downloader.Root(cmd, args)
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	downloadCmd.Flags().StringP("game", "g", "", "Specify a game to download for.")
}
