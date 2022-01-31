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
	"os"

	"github.com/spf13/cobra"
)

var jsonInput string
var schemeInput string
var userInput string
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
	Long: `Build a URI or URL from its components.

Examples:

	url build --scheme http --host myhost.com
		http://myhost.com

	url build --scheme mailto --uri-path myemail@myhost.com
		mailto:myemail@myhost.com

	url build --scheme http --host myhost.com --port 8888 --path /colorado/denver
		http://myhost.com:8888/colorado/denver

	url build --scheme http --host myhost.com --param foo=bar --param bar=baz
		http://myhost.com?bar=baz&foo=bar

	url build --scheme http --host myhost.com --query "foo=bar&bar=baz"
		http://myhost.com?foo=bar&bar=baz

	url build --json '{"scheme":"http","host":"myhost.com","params":{"foo":"bar","bar":"baz"}}'
		http://myhost.com?bar=baz&foo=bar
	`,
	Args: cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var uri URI
		if jsonInput != "" {
			uri = buildURIFromJSON(jsonInput)
		} else {
			uri = URI{
				Scheme:    schemeInput,
				User:      userInput,
				UriPath:   uriPathInput,
				Host:      domainInput,
				Port:      portInput,
				Path:      pathInput,
				Fragment:  fragmentInput,
				Query:     queryInput,
				RawParams: paramsInput,
			}
		}
		u := uri.AsURL()
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
	buildCmd.Flags().StringVar(&userInput, userLabel, "", "Provides a user[:password].")
	buildCmd.Flags().StringVar(&domainInput, hostLabel, "", "Provide a URI authority/domain/host or host:port.")
	buildCmd.Flags().StringVar(&portInput, portLabel, "", "Provide a port number.")
	buildCmd.Flags().StringVar(&pathInput, pathLabel, "", "Provide a URL path.")
	buildCmd.Flags().StringVar(&uriPathInput, "uri-path", "", "Provides URI (not URL) path.")
	buildCmd.Flags().StringVar(&fragmentInput, fragmentLabel, "", "Provide a URI fragment.")
	buildCmd.Flags().StringVar(&queryInput, "query", "", "Provide a URL query string (without ?).")
	buildCmd.Flags().StringArrayVar(&paramsInput, paramLabel, nil, "Provide a key=value pair of query parameters.")
}
