package pkg

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

/*
URL Encode & Decode Data

*/

// Flags related to url encoding and decoding
var (
	urldecode = false // URL decode
)

// urlEnc : Url encode and decode module
var urlEnc = &cobra.Command{
	Use:   "url",
	Short: "A Simple URL Encoder and Decoder",
	Long:  "A Simple URL Encoder & Decoder with Similar syntax to base64 unix command",
	Example: `
// URL Encode data
echo "some Url encoded_data" | bninja url

// URL Decode data
echo "some+Url+encoded%5\fdata" | bninja url -d
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		input := synchronousInput(args)

		if urldecode {
			result, er := url.QueryUnescape(input)
			if er != nil {
				return fmt.Errorf("failed to url decode data %v", er)
			}
			manageOutput(result)
		} else {
			manageOutput(url.QueryEscape(input))
		}

		return nil
	},
}

func init() {
	urlEnc.Flags().BoolVarP(&urldecode, "decode", "d", false, "URL Decode")

	urlEnc.CompletionOptions.DisableDefaultCmd = true
}
