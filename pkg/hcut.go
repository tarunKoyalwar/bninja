package pkg

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// hcut : cut data horizantally
var hcut = &cobra.Command{
	Use:   "hcut",
	Short: "Cut Data horizantally similar to GNU cut but horizantal",
	Long:  "Cut Data horizantally similar to GNU cut but horizantal",
	Example: `
// cut data horizantally
cat somefile.txt | bninja hcut -d "delimeter" -f 1,7-8
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		input := synchronousInput(args)

		getindexes()

		if len(index) == 0 || delimeter == "" {
			return fmt.Errorf("fields or delimeter was not given of failed to extract")
		}

		data := strings.Split(input, delimeter)

		finalarr := []string{}

		for _, i := range index {
			if i-1 >= len(data) {
				continue
			}
			finalarr = append(finalarr, data[i-1])
		}
		manageOutput(strings.Join(finalarr, " "))

		return nil
	},
}

func init() {
	hcut.Flags().StringVarP(&delimeter, "delim", "d", "", "Delimeter to Use(Supports Multiple Chars)")
	hcut.Flags().StringVarP(&fields, "fields", "f", "", "fields/colums but it can be 1-3 or 1,5,6 or 1-3,8 (Inclusive)(starts from 1)")

	_ = hcut.MarkFlagRequired(delimeter)
	_ = hcut.MarkFlagRequired(fields)

	hcut.CompletionOptions.DisableDefaultCmd = true
}
