/*
Copyright © 2022 Lloyd Lobo <hello@lloydlobo.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/lloydlobo/okejoke/api"
	"github.com/lloydlobo/okejoke/logo"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `OkeJoke searches for list of jokes that match your search term.
        _         _       _        
   ___ | | _____ (_) ___ | | _____ 
  / _ \| |/ / _ \| |/ _ \| |/ / _ \
 | (_) |   <  __/| | (_) |   <  __/
  \___/|_|\_\___|/ |\___/|_|\_\___|
               |__/                
    `,
	Run: func(cmd *cobra.Command, args []string) {
		logo.PrintLogo()
		fmt.Println("search called")
		api.GetOkejoke() // Basic "random" command.
		// TODO: Add flags `page` 'limit' 'term'
		/*
			GET https://icanhazdadjoke.com/search - search for dad jokes.
			This endpoint accepts the following optional query string parameters:
			page - which page of results to fetch (default: 1)
			limit - number of results to return per page (default: 20) (max: 30)
			term - search term to use (default: list all jokes)
		*/
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
