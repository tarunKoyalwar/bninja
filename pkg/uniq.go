package pkg

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tarunKoyalwar/bninja/pkg/internal"
)

/*
Get / Print Unique items
*/

var (
	insensitive   = false
	showfrequency = false
)

// unqiue : get unique lines
var unique = &cobra.Command{
	Use:   "uniq",
	Short: "Print Unique lines",
	Long:  "Print Unique Lines (uses golang map to filter duplicates)",
	Example: `
// Print Unique Lines Only
cat subdomains.txt | bninja uniq
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		input := synchronousInput(args)

		keys := map[string]int{}

		for _, v := range internal.Split(input, '\n') {
			if insensitive {
				v = strings.ToLower(v)
			}
			x := strings.TrimSpace(v)
			keys[x] += 1
		}

		var buffer bytes.Buffer

		for k, v := range keys {
			if showfrequency {
				buffer.WriteString(k + ":" + strconv.Itoa(v) + "\n")
			} else {
				buffer.WriteString(k + "\n")
			}
		}

		manageOutput(buffer.String())

		return nil
	},
}

func init() {
	unique.Flags().BoolVar(&insensitive, "insensitive", false, "ignore case/ insensitive")
	unique.Flags().BoolVar(&showfrequency, "show-frequency", false, "Show frequency of each item(i.e no of times it appeared")

	unique.CompletionOptions.DisableDefaultCmd = true
}
