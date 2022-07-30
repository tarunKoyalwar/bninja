package pkg

import (
	"encoding/json"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
)

var (
	sorted    = false //sort keys
	jsoncolor = false
)

// jsonn : json module to format and colorize
var jsonn = &cobra.Command{
	Use:   "json",
	Short: "Format JSON",
	Long:  "Format JSON and Color(Linux Only) returns given data if invalid json",
	Example: `
// Format JSON and colorize it
curl -v http://somewhere/data.json | bninja json --color

// Sort Keys In Ascending Order
cat data.json | bninja json --sort
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		input := synchronousInput(args)

		var z map[string]any

		err := json.Unmarshal([]byte(input), &z)
		if err == nil {
			//valid json
			results := pretty.PrettyOptions([]byte(input), &pretty.Options{
				Indent: "	",
				SortKeys: sorted,
			})

			if jsoncolor && runtime.GOOS == "linux" {
				final := pretty.Color(results, nil)
				manageOutput(string(final))
			} else {
				manageOutput(string(results))
			}

		} else {
			manageOutput(input)
		}

		return nil
	},
}

func init() {
	jsonn.Flags().BoolVar(&sorted, "sort", false, "Sort JSON Keys")
	jsonn.Flags().BoolVar(&jsoncolor, "color", false, "Colorize JSON (Only Linux)")

	jsonn.CompletionOptions.DisableDefaultCmd = true
}
