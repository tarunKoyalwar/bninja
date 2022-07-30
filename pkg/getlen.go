package pkg

import (
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// getlenc : get length of given string
var getlenc = &cobra.Command{
	Use:   "getlen",
	Short: "Returns length of the given string",
	Long:  "Returns length of the given string",
	Example: `
// get length of a string
echo "some string" | bninja getlen
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		input := synchronousInput(args)
		result := len(strings.TrimSpace(input))
		manageOutput(strconv.Itoa(result))

		return nil
	},
}

func init() {
	getlenc.CompletionOptions.DisableDefaultCmd = true
}
