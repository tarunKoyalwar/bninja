package pkg

import (
	"bytes"
	"fmt"
	"os"
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
	summary       = false
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
		lines := internal.Split(input, '\n')

		keys := map[string]int{}

		for _, v := range lines {
			x := strings.TrimSpace(v)
			if x == "" {
				continue
			}
			if insensitive {
				x = strings.ToLower(v)
			}
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

		if summary {
			fmt.Fprintf(os.Stderr, "Input Lines: %v, Unique Lines: %v , Duplicates: %v\n", len(lines), len(keys), (len(lines) - len(keys)))
		}
		return nil
	},
}

func init() {
	unique.Flags().BoolVar(&insensitive, "insensitive", false, "ignore case/ insensitive")
	unique.Flags().BoolVar(&showfrequency, "show-frequency", false, "Show frequency of each item(i.e no of times it appeared")
	unique.Flags().BoolVar(&summary, "summary", false, "Shows stats related to input and unique lines")
	unique.CompletionOptions.DisableDefaultCmd = true
}
