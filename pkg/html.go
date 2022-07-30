package pkg

import (
	"html"

	"github.com/spf13/cobra"
)

/*
HTML Encoding & Decoding

*/

var (
	htmldecode = false
)

// htmlenc : html enc and decode module
var htmlenc = &cobra.Command{
	Use:   "html",
	Short: "HTML Encoder/Decoder",
	Long:  "Encode and Decode HTML",
	Example: `
// HTML Encode data
echo "Fran & Freddie's Diner" <tasty@example.com>" | bninja html

// HTML Decode data
echo "Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;" | bninja html -d
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		input := synchronousInput(args)

		if urldecode {
			manageOutput(html.UnescapeString(input))
		} else {
			manageOutput(html.EscapeString(input))
		}

		return nil
	},
}

func init() {
	htmlenc.Flags().BoolVarP(&htmldecode, "decode", "d", false, "HTML Decode")

	htmlenc.CompletionOptions.DisableDefaultCmd = true
}
