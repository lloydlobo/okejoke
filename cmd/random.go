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
	"github.com/spf13/cobra"

	"github.com/lloydlobo/okejoke/api"
	"github.com/lloydlobo/okejoke/logo"
	// "github.com/lloydlobo/okejoke/spin"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:        "random",
	Aliases:    []string{"rand", "r"},
	SuggestFor: []string{"rando", "rnd"},
	Short:      "Get a random okejoke",
	Long: `OkeJoke CLI generates jokes on the fly. Oke? Bye!
        _         _       _        
   ___ | | _____ (_) ___ | | _____ 
  / _ \| |/ / _ \| |/ _ \| |/ / _ \
 | (_) |   <  __/| | (_) |   <  __/
  \___/|_|\_\___|/ |\___/|_|\_\___|
               |__/                
    `,
	Run: func(cmd *cobra.Command, args []string) {
		logo.PrintLogo()
		// go spin.RenderSpinner() // FIXME: spinner needs to exit as it hides the nblinking cursor.

		// Flags returns the complete FlagSet that applies
		// to this command \(local and persistent declared here and by all parents\)\.
		jokeTerm, _ := cmd.Flags().GetString("term") // GetString return the string value of a flag with the given name
		emptyFlag := ""

		if jokeTerm != emptyFlag {
			// cmd with flag "term", For Example; `--term=hello`.
			api.GetRandomJokeWithTerm(jokeTerm)
			api.FetchUserTermFlagData(jokeTerm)
		} else {
			api.GetOkejoke() // Basic "random" command.
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	// Use this flag in function GetRandomJokesWIthTerm.
	// It runs every time user types a flag with random cmd.
	randomCmd.
		PersistentFlags().                                 // PersistentFlags returns the persistent FlagSet specifically set in the current command.
		String("term", "", "A search term for a okejoke.") // String defines a string flag with specified name, default value, and usage string\. The return value is the address of a string variable that stores the value of the flag\.
}

// https://youtu.be/kT7Z02bR1IY

/*
    Example: "",
	ValidArgs: []string{},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	},
	Args: func(cmd *cobra.Command, args []string) error {
	},
	ArgAliases:             []string{},
	BashCompletionFunction: "",
	Deprecated:             "",
	Annotations:            map[string]string{},
	Version:                "",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
	},
	PreRun: func(cmd *cobra.Command, args []string) {
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
	},
	RunE: func(cmd *cobra.Command, args []string) error {
	},
	PostRun: func(cmd *cobra.Command, args []string) {
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
	},
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
	CompletionOptions:          cobra.CompletionOptions{},
	TraverseChildren:           false,
	Hidden:                     false,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
*/
