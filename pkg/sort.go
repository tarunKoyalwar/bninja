package pkg

import (
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tarunKoyalwar/bninja/pkg/internal"
)

/*
sort and dump given data

*/
var reverse bool

// sortc : sort module
var sortc = &cobra.Command{
	Use:   "sort",
	Short: "Sort Input data in ascending/descending order",
	Long:  "Sort Input data in ascending/descending order. Will trim spaces by default",
	Example: `
// sort data in ascending order
cat strings | bninja sort
`,
	Run: func(cmd *cobra.Command, args []string) {

		data := synchronousInput(args)
		tarr := internal.Split(data, '\n')

		tdata := []string{}

		for _, v := range tarr {
			x := strings.TrimSpace(v)
			if v != "" {
				tdata = append(tdata, x)
			}
		}

		sort.Sort(sort.StringSlice(tdata))
		if !reverse {
			manageOutput(strings.Join(tdata, "\n"))
		} else {
			sort.Sort(sort.Reverse(sort.StringSlice(tdata)))
			manageOutput(strings.Join(tdata, "\n"))
		}
	},
}

func init() {
	sortc.Flags().BoolVarP(&reverse, "reverse", "r", false, "Sort data in descending Order")

	sortc.CompletionOptions.DisableDefaultCmd = true
}
