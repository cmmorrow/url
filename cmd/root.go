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
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const schemeLabel = "scheme"
const opaqueLabel = "uri-path"
const userLabel = "user"
const hostLabel = "host"
const portLabel = "port"
const pathLabel = "path"
const fragmentLabel = "fragment"
const paramsLabel = "params"
const paramLabel = "param"

var puny bool
var shell bool

type URI struct {
	Scheme    string
	UriPath   string
	User      string
	Host      string
	Port      string
	Path      string
	Fragment  string
	Query     string
	RawParams []string
	Params    map[string]interface{}
}

func (u *URI) AsURL() url.URL {
	var output = url.URL{
		Scheme:   u.Scheme,
		Opaque:   u.UriPath,
		Host:     u.BuildHostname(),
		Path:     u.Path,
		Fragment: u.Fragment,
	}
	if u.User != "" {
		output.User = u.SetUser()
	}
	if u.Query != "" {
		output.RawQuery = u.Query
	} else {
		output.RawQuery = u.BuildValues().Encode()
	}
	return output
}

func (u *URI) SetUser() *url.Userinfo {
	var username, password string
	if u.User != "" {
		user := strings.Split(u.User, ":")
		if len(user) == 1 {
			user = append(user, "")
		}
		username, password = user[0], user[1]
	}
	if password != "" {
		return url.UserPassword(username, password)
	} else {
		return url.User(username)
	}
}

func (u *URI) BuildHostname() string {
	if u.Host == "" {
		return ""
	}
	if u.Port != "" {
		splitHost := strings.SplitN(u.Host, ":", 2)
		return splitHost[0] + ":" + u.Port
	} else {
		return u.Host
	}
}

func (u *URI) BuildValues() url.Values {
	values := url.Values{}
	switch {
	case u.Params != nil:

		for key, val := range u.Params {
			switch v := val.(type) {
			case string:
				values[key] = []string{v}
			case []interface{}:
				var vv []string
				for i := range v {
					vv = append(vv, v[i].(string))
				}
				values[key] = vv
			}
		}
		return values
	case u.RawParams != nil:
		for i := range u.RawParams {
			p := strings.SplitN(u.RawParams[i], "=", 2)
			if len(p) == 1 {
				p = append(p, "")
			}
			values[p[0]] = append(values[p[0]], p[1])
		}
	}

	return values
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "url",
	Short: "A command-line tool for working with URLs.",
	Long: `

	_   _ ___ _      _____ ___   ___  _    
	| | | | _ \ |    |_   _/ _ \ / _ \| |   
	| |_| |   / |__    | || (_) | (_) | |__ 
	 \___/|_|_\____|   |_| \___/ \___/|____|
											
   

	A command-line tool for working with URLs.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&puny, "puny", false, "Convert the domain/host to punycode (IDNA).")
	rootCmd.PersistentFlags().BoolVar(&shell, "shell", false, "Remove shell escape characters before processing.")
}
