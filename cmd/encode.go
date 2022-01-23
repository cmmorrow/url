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
	"golang.org/x/net/idna"
)

var puny bool

// encodeCmd represents the encode command
var encodeCmd = &cobra.Command{
	Use:   "encode string",
	Short: "Encode a URL.",
	Long:  `Percent encode a string into valid URL.`,
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		var input string = args[0]
		url, err := url.Parse(input)
		if err != nil {
			fmt.Printf("Error parsing %s\n", input)
			os.Exit(1)
		}
		if puny {
			var puny *idna.Profile = idna.New()
			out, err := puny.ToASCII(input)
			if err != nil {
				fmt.Printf("Error coverting %s to punycode.", input)
			}
			fmt.Printf("%s\n", out)
			return
		}
		fmt.Printf("%s\n", url.String())
	},
}

func init() {
	rootCmd.AddCommand(encodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	encodeCmd.Flags().BoolVar(&puny, "puny", false, "Convert the domain/host to punycode (IDNA).")
}
