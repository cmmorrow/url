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

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse string",
	Short: "Parse a URL into its components.",
	Long: `Parse a URL into its primary components.
	
	Each matching component is displayed on a new line.`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(1), cobra.MinimumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		var input string = args[0]
		url, err := url.Parse(input)
		if err != nil {
			fmt.Printf("Error parsing %s\n", input)
			os.Exit(1)
		}

		fmt.Printf("scheme:	%s\n", url.Scheme)
		fmt.Printf("opaque:	%s\n", url.Opaque)
		fmt.Printf("domain:	%s\n", url.Hostname())
		fmt.Printf("port:	%s\n", url.Port())
		fmt.Printf("path:	%s\n", url.Path)
		fmt.Printf("fragment:	%s\n", url.Fragment)

		for key, vals := range url.Query() {
			for i := range vals {
				fmt.Printf("param:	%s=%s\n", key, vals[i])
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
