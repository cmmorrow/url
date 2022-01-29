/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var jsonInput string
var schemeInput string
var domainInput string
var portInput string
var pathInput string
var uriPathInput string
var fragmentInput string
var queryInput string
var paramsInput []string

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a new URI.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var uri URI
		if jsonInput != "" {
			uri = buildURIFromJSON(jsonInput)
		} else {
			uri = URI{
				Scheme:   schemeInput,
				Opaque:   uriPathInput,
				Domain:   domainInput,
				Port:     portInput,
				Path:     pathInput,
				Fragment: fragmentInput,
				Query:    queryInput,
				Params:   paramsInput,
			}
		}
		u := uri.asURL()
		fmt.Printf("%s\n", u.String())
	},
}

func buildURIFromJSON(j string) URI {
	var uri = URI{}
	err := json.Unmarshal([]byte(j), &uri)
	if err != nil {
		fmt.Println("Error reading JSON.")
		os.Exit(1)
	}
	return uri
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().StringVar(&jsonInput, "json", "", "Provide input as JSON.")
	buildCmd.Flags().StringVar(&schemeInput, schemeLabel, "", "Provide a URI scheme (or protocol).")
	buildCmd.Flags().StringVar(&domainInput, domainLabel, "", "Provide a URI authority/domain/host or host:port.")
	buildCmd.Flags().StringVar(&portInput, portLabel, "", "Provide a port number.")
	buildCmd.Flags().StringVar(&pathInput, pathLabel, "", "Provide a URL path.")
	buildCmd.Flags().StringVar(&uriPathInput, "uri-path", "", "Provides URI (not URL) path.")
	buildCmd.Flags().StringVar(&fragmentInput, fragmentLabel, "", "Provide a URI fragment.")
	buildCmd.Flags().StringVar(&queryInput, "query", "", "Provide a URL query string (without ?).")
	buildCmd.Flags().StringArrayVar(&paramsInput, paramLabel, nil, "Provide a key=value pair of query parameters.")
}
