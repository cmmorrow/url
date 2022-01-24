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
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/net/idna"
)

var scheme bool
var opaque bool
var domain bool
var port bool
var path bool
var fragment bool
var params bool
var noColor bool
var jsonOutput bool
var noDecode bool

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
		case scheme:
			displayComponent("", url.Scheme)
		case opaque:
			displayComponent("", url.Opaque)
		case domain:
			displayComponent("", hostname(url))
		case port:
			displayComponent("", url.Port())
		case path:
			if noDecode {
				displayComponent("", url.Path)
			} else {
				displayComponent("", url.EscapedPath())
			}
		case fragment:
			if noDecode {
				displayComponent("", url.Fragment)
			} else {
				displayComponent("", url.EscapedFragment())
			}
		case params:
			for key, vals := range url.Query() {
				for i := range vals {
					displayComponent("", key+"="+vals[i])
				}
			}
		default:
			if jsonOutput {
				jsonData := make(map[string]interface{})

				jsonData["scheme"] = url.Scheme
				if jsonData["scheme"] == "" {
					jsonData["scheme"] = nil
				}
				jsonData["opaque"] = url.Opaque
				if jsonData["opaque"] == "" {
					jsonData["opaque"] = nil
				}
				jsonData["domain"] = hostname(url)
				if jsonData["domain"] == "" {
					jsonData["domain"] = nil
				}
				jsonData["port"] = url.Port()
				if jsonData["port"] == "" {
					jsonData["port"] = nil
				}
				if noDecode {
					jsonData["path"] = url.Path
					if jsonData["path"] == "" {
						jsonData["path"] = nil
					}
					jsonData["fragment"] = url.Fragment
					if jsonData["fragment"] == "" {
						jsonData["fragment"] = nil
					}
				} else {
					jsonData["path"] = url.EscapedPath()
					if jsonData["path"] == "" {
						jsonData["path"] = nil
					}
					jsonData["fragment"] = url.EscapedFragment()
					if jsonData["fragment"] == "" {
						jsonData["fragment"] = nil
					}
				}

				q := url.Query()
				queryParameters := make(map[string]interface{})
				if len(q) > 0 {
					for key, vals := range q {
						if len(vals) > 1 {
							queryParameters[key] = vals
						} else {
							queryParameters[key] = vals[0]
						}
					}
				}
				jsonData["params"] = queryParameters
				b, err := json.Marshal(jsonData)
				if err != nil {
					fmt.Println("Cannot convert to JSON.")
					os.Exit(1)
				}
				fmt.Println(string(b))
				return
			}

			displayComponent("scheme", url.Scheme)
			displayComponent("opaque", url.Opaque)
			displayComponent("domain", hostname(url))
			displayComponent("port", url.Port())
			if noDecode {
				displayComponent("path", url.Path)
				displayComponent("fragment", url.Fragment)
			} else {
				displayComponent("path", url.EscapedPath())
				displayComponent("fragment", url.EscapedFragment())
			}

			q := url.Query()
			if len(q) == 0 {
				displayComponent("param", "")
				return
			}
			for key, vals := range q {
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
	if noColor {
		fmt.Printf("%s: %s\n", displayName, value)
		return
	}
	blue := color.New(color.Bold, color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("%s:	%s\n", blue(displayName), green(value))
}

func hostname(u *url.URL) string {
	if puny {
		var p *idna.Profile = idna.New()
		host := u.Hostname()
		out, err := p.ToASCII(host)
		if err != nil {
			fmt.Printf("Error coverting %s to punycode.\n", host)
			os.Exit(1)
		}
		return out
	} else {
		return u.Hostname()
	}
}

func init() {
	rootCmd.AddCommand(parseCmd)

	parseCmd.Flags().BoolVar(&scheme, "scheme", false, "Only display the scheme.")
	parseCmd.Flags().BoolVar(&opaque, "opaque", false, "Only display the opaque.")
	parseCmd.Flags().BoolVar(&domain, "domain", false, "Only display the domain/host.")
	parseCmd.Flags().BoolVar(&port, "port", false, "Only display the port number.")
	parseCmd.Flags().BoolVar(&path, "path", false, "Only display the path.")
	parseCmd.Flags().BoolVar(&fragment, "fragment", false, "Only display the URL fragment.")
	parseCmd.Flags().BoolVar(&params, "params", false, "Only display the query parameters.")
	parseCmd.Flags().BoolVar(&noColor, "no-color", false, "Suppress color text output.")
	parseCmd.Flags().BoolVar(&jsonOutput, "json", false, "Output as JSON.")
	parseCmd.Flags().BoolVar(&noDecode, "no-decode", false, "Do not URL decode paths and query parameters.")
}
