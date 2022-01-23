/*
Copyright © 2022 Chris Morrow cmmorrow@gmail.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var Scheme bool
var Opaque bool
var Domain bool
var Port bool
var Path bool
var Fragment bool
var Params bool

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

		switch {
		case Scheme:
			displayComponent("", url.Scheme)
		case Opaque:
			displayComponent("", url.Opaque)
		case Domain:
			displayComponent("", url.Hostname())
		case Port:
			displayComponent("", url.Port())
		case Path:
			displayComponent("", url.Path)
		case Fragment:
			displayComponent("", url.Fragment)
		case Params:
			for key, vals := range url.Query() {
				for i := range vals {
					displayComponent("", key+"="+vals[i])
				}
			}
		default:
			displayComponent("scheme", url.Scheme)
			displayComponent("opaque", url.Opaque)
			displayComponent("domain", url.Hostname())
			displayComponent("port", url.Port())
			displayComponent("path", url.Path)
			displayComponent("fragment", url.Fragment)

			for key, vals := range url.Query() {
				for i := range vals {
					displayComponent("param", key+"="+vals[i])
				}
			}
		}
	},
}

func displayComponent(displayName string, value string) {
	if displayName == "" {
		fmt.Printf("%s\n", value)
		return
	}
	fmt.Printf("%s:	%s\n", displayName, value)
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
	parseCmd.Flags().BoolVar(&Scheme, "scheme", false, "Only display the scheme.")
	parseCmd.Flags().BoolVar(&Opaque, "opaque", false, "Only display the opaque.")
	parseCmd.Flags().BoolVar(&Domain, "domain", false, "Only display the domain/host.")
	parseCmd.Flags().BoolVar(&Port, "port", false, "Only display the port number.")
	parseCmd.Flags().BoolVar(&Path, "path", false, "Only display the path.")
	parseCmd.Flags().BoolVar(&Fragment, "fragment", false, "Only display the URL fragment.")
	parseCmd.Flags().BoolVar(&Params, "params", false, "Only display the query parameters.")
}
