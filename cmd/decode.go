/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode string",
	Short: "Decode a URL.",
	Long:  `Decode a URL encoded string.`,
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		var input string = args[0]
		url, err := url.PathUnescape(input)
		if err != nil {
			fmt.Printf("Error parsing %s\n", input)
			os.Exit(1)
		}
		fmt.Printf("%s\n", url)
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
