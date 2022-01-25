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
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/net/idna"
)

var schemeFlag bool
var opaqueFlag bool
var domainFlag bool
var portFlag bool
var pathFlag bool
var fragmentFlag bool
var paramsFlag bool

var noColorFlag bool
var jsonOutputFlag bool
var noDecodeFlag bool

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse string",
	Short: "Parse a URL into its components.",
	Long: `Parse a URL into its primary components.
	
	Each matching component is displayed on a new line.`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(1), cobra.MinimumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		var input string = args[0]
		if shell {
			input = strings.ReplaceAll(input, "\\", "")
		}
		urlStruct, err := url.Parse(input)
		if err != nil {
			fmt.Printf("Error parsing %s\n", input)
			os.Exit(1)
		}

		switch {
		case schemeFlag:
			displayComponent("", urlStruct.Scheme)
		case opaqueFlag:
			displayComponent("", urlStruct.Opaque)
		case domainFlag:
			displayComponent("", hostname(urlStruct))
		case portFlag:
			displayComponent("", urlStruct.Port())
		case pathFlag:
			if noDecodeFlag {
				displayComponent("", urlStruct.EscapedPath())
			} else {
				displayComponent("", urlStruct.Path)
			}
		case fragmentFlag:
			if noDecodeFlag {
				displayComponent("", urlStruct.EscapedFragment())
			} else {
				displayComponent("", urlStruct.Fragment)
			}
		case paramsFlag:
			for key, vals := range urlStruct.Query() {
				for i := range vals {
					displayComponent("", key+"="+vals[i])
				}
			}
		default:
			if jsonOutputFlag {
				b := buildJSON(urlStruct)
				fmt.Println(b)
				return
			}

			displayComponent(schemeLabel, urlStruct.Scheme)
			displayComponent(opaqueLabel, urlStruct.Opaque)
			displayComponent(domainLabel, hostname(urlStruct))
			displayComponent(portLabel, urlStruct.Port())
			if noDecodeFlag {
				displayComponent(pathLabel, urlStruct.EscapedPath())
				displayComponent(fragmentLabel, urlStruct.EscapedFragment())
			} else {
				displayComponent(pathLabel, urlStruct.Path)
				displayComponent(fragmentLabel, urlStruct.Fragment)
			}

			q := urlStruct.Query()
			if len(q) == 0 {
				displayComponent(paramLabel, "")
				return
			}
			for key, vals := range q {
				for i := range vals {
					displayComponent(paramLabel, key+"="+vals[i])
				}
			}
		}
	},
}

func buildJSON(u *url.URL) string {
	jsonData := make(map[string]interface{})

	jsonData[schemeLabel] = u.Scheme
	if jsonData[schemeLabel] == "" {
		jsonData[schemeLabel] = nil
	}
	jsonData[opaqueLabel] = u.Opaque
	if jsonData[opaqueLabel] == "" {
		jsonData[opaqueLabel] = nil
	}
	jsonData[domainLabel] = hostname(u)
	if jsonData[domainLabel] == "" {
		jsonData[domainLabel] = nil
	}
	jsonData[portLabel] = u.Port()
	if jsonData[portLabel] == "" {
		jsonData[portLabel] = nil
	}
	if noDecodeFlag {
		jsonData[pathLabel] = u.EscapedPath()
		if jsonData[pathLabel] == "" {
			jsonData[pathLabel] = nil
		}
		jsonData[fragmentLabel] = u.EscapedFragment()
		if jsonData[fragmentLabel] == "" {
			jsonData[fragmentLabel] = nil
		}
	} else {
		jsonData[pathLabel] = u.Path
		if jsonData[pathLabel] == "" {
			jsonData[pathLabel] = nil
		}
		jsonData[fragmentLabel] = u.Fragment
		if jsonData[fragmentLabel] == "" {
			jsonData[fragmentLabel] = nil
		}
	}

	q := u.Query()
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
	jsonData[paramsLabel] = queryParameters
	b, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Cannot convert to JSON.")
		os.Exit(1)
	}
	return string(b)
}

func displayComponent(displayName string, value string) {
	if displayName == "" {
		fmt.Printf("%s\n", value)
		return
	}
	if noColorFlag {
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

	parseCmd.Flags().BoolVar(&schemeFlag, schemeLabel, false, "Only display the scheme.")
	parseCmd.Flags().BoolVar(&opaqueFlag, opaqueLabel, false, "Only display the opaque.")
	parseCmd.Flags().BoolVar(&domainFlag, domainLabel, false, "Only display the domain/host.")
	parseCmd.Flags().BoolVar(&portFlag, portLabel, false, "Only display the port number.")
	parseCmd.Flags().BoolVar(&pathFlag, pathLabel, false, "Only display the path.")
	parseCmd.Flags().BoolVar(&fragmentFlag, fragmentLabel, false, "Only display the URL fragment.")
	parseCmd.Flags().BoolVar(&paramsFlag, paramsLabel, false, "Only display the query parameters.")
	parseCmd.Flags().BoolVar(&noColorFlag, "no-color", false, "Suppress color text output.")
	parseCmd.Flags().BoolVar(&jsonOutputFlag, "json", false, "Output as JSON.")
	parseCmd.Flags().BoolVar(&noDecodeFlag, "no-decode", false, "Do not URL decode paths and query parameters.")
}
