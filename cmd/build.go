/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

var jsonInput string
var schemeInput string
var domainInput string
var portInput string
var pathInput string
var uriPathInput string
var fragmentInput string
var paramsInput []string

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if len(portInput) > 0 {
			host := strings.SplitN(domainInput, ":", 2)
			domainInput = host[0] + ":" + portInput
		}

		values := url.Values{}
		for i := range paramsInput {
			params := strings.SplitN(paramsInput[i], "=", 2)
			if len(params) != 2 {
				params[1] = ""
			}
			values[params[0]] = append(values[params[0]], params[1])
		}
		urlStruct := url.URL{
			Scheme:   schemeInput,
			Opaque:   uriPathInput,
			Host:     domainInput,
			Path:     pathInput,
			RawQuery: values.Encode(),
			Fragment: fragmentInput,
		}
		fmt.Printf("%s\n", urlStruct.String())
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	buildCmd.Flags().StringVar(&jsonInput, "json", "", "Provide input as JSON.")
	buildCmd.Flags().StringVar(&schemeInput, schemeLabel, "", "Provide a URI scheme (or protocol).")
	buildCmd.Flags().StringVar(&domainInput, domainLabel, "", "Provide a URI authority/domain/host or host:port.")
	buildCmd.Flags().StringVar(&portInput, portLabel, "", "Provide a port number.")
	buildCmd.Flags().StringVar(&pathInput, pathLabel, "", "Provide a URL path.")
	buildCmd.Flags().StringVar(&uriPathInput, "uri-path", "", "Provides URI (not URL) path.")
	buildCmd.Flags().StringVar(&fragmentInput, fragmentLabel, "", "Provide a URI fragment.")
	buildCmd.Flags().StringArrayVar(&paramsInput, paramLabel, nil, "Provide a key=value pair of query parameters.")
}
