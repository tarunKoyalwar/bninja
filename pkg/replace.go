package pkg

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

/*
match and replace without being eyesore

*/

var (
	count = -1    // replace all occurences
	regex = false // regex mode
	m     = ""    // string to match
	r     = ""
	posix = false
)

// replace : match and replace data
var replace = &cobra.Command{
	Use:   "replace",
	Short: "Match & Replace Data",
	Long:  "Match & Replace Data from files In Regex Mode Go(RE2) Engine is used",
	Example: `
// replace first 5 occurences of "programmer" to "dev"
bninja replace -m "programmer" -r "dev" -c 5
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if m == "" {
			return fmt.Errorf("match string not given")
		}

		m = InterpretString(m)
		r = InterpretString(r)

		input := synchronousInput(args)

		if !regex {
			data := strings.Replace(input, m, r, count)
			manageOutput(data)
		} else {
			var rxp *regexp.Regexp
			if posix {
				rxp = regexp.MustCompilePOSIX(m)
			} else {
				rxp = regexp.MustCompile(m)
			}

			if count == -1 {
				manageOutput(rxp.ReplaceAllString(input, r))
			} else {
				indexes := rxp.FindAllIndex([]byte(input), count)
				if indexes == nil {
					os.Exit(1) //i.e no match found
				} else {
					if len(indexes) == count {
						manageOutput(rxp.ReplaceAllString(input, r))
					} else {
						var buff bytes.Buffer
						counter := 0
						for _, v := range indexes {
							buff.WriteString(input[counter:v[0]])
							buff.WriteString(r)
							counter = v[1]
						}
						buff.WriteString(input[counter:])

						manageOutput(buff.String())
					}
				}
			}
		}

		return nil
	},
}

func init() {
	replace.Flags().IntVarP(&count, "count", "c", -1, "No of Occurences to replace(default -1 i.e all)")
	replace.Flags().BoolVar(&regex, "regex", false, "Regex Mode")
	replace.Flags().StringVarP(&m, "match", "m", "", "String to Match")
	replace.Flags().StringVarP(&r, "replace", "r", "", "replace matched character with string (can be empty)")
	replace.Flags().BoolVar(&posix, "posix", false, "Use POSIX for regex")

	replace.CompletionOptions.DisableDefaultCmd = true
}

// Converts a raw string to interpreted string
func InterpretString(raw string) string {
	quoted := strconv.Quote(raw)
	quoted = strings.ReplaceAll(quoted, `\\`, `\`)
	if unquoted, err := strconv.Unquote(quoted); err == nil {
		return unquoted
	}
	return raw
}
