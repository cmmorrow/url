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
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/idna"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode string",
	Short: "Decode a URL.",
	Long:  `Decode a URL encoded string.`,
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		var input string = args[0]
		if shell {
			input = strings.ReplaceAll(input, "\\", "")
		}
		if puny {
			var p *idna.Profile = idna.New()
			out, err := p.ToUnicode(input)
			if err != nil {
				fmt.Printf("Error converting %s.\n", input)
				os.Exit(1)
			}
			fmt.Printf("%s\n", out)
		} else {
			decoded, err := url.QueryUnescape(input)
			if err != nil {
				fmt.Printf("Error decoding %s.\n", input)
				os.Exit(1)
			}
			fmt.Printf("%s\n", decoded)
		}
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}
