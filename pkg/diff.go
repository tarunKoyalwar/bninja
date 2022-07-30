package pkg

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/tarunKoyalwar/bninja/pkg/internal"
)

var (
	file1  = false // only print file1 unique items
	file2  = false // only print file2 unique items
	common = false // only print common unique lines
)

// diffc provides diff command
var diffc = &cobra.Command{
	Use:   "diff",
	Short: "Get unique and Common lines of files",
	Long: `Get Unique or Common Lines of Two Files.
It doesnot matter if it has extra spaces or lines are shuffled/proportional 
Implementation uses Golang Maps to get Common and unique lines.
	`,
	Example: `
// Common Unique Lines
bninja diff file1 file2 -c

// Unique Lines of file1
bninja diff file1 file2 -1

//Unique Lines of file2
bninja diff file1 file2 -2
`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		flagtest := file1 || file2 || common
		if !flagtest {
			panic("At least One Flag requireds")
		}

		data1 := internal.ReadFromFile(args[0])
		data2 := internal.ReadFromFile(args[1])

		map1 := map[string]bool{}
		map2 := map[string]bool{}

		uniqf1lines := []string{}
		uniqf2lines := []string{}
		commlines := []string{}

		//map all file1 lines to keys
		for _, v := range internal.Split(data1, '\n') {
			v = strings.TrimSpace(v)
			if v == "" {
				continue
			}
			map1[v] = true
		}
		// map all file2 lines to keys
		for _, v := range internal.Split(data2, '\n') {
			v = strings.TrimSpace(v)
			if v == "" {
				continue
			}
			map2[v] = true
		}

		if common {
			for k := range map1 {
				if map2[k] {
					commlines = append(commlines, k)
				}
			}
			manageOutput(strings.Join(commlines, "\n"))
		} else if file1 {
			for k := range map1 {
				if !map2[k] {
					uniqf1lines = append(uniqf1lines, k)
				}
			}
			manageOutput(strings.Join(uniqf1lines, "\n"))
		} else if file2 {
			for k := range map2 {
				if !map1[k] {
					uniqf2lines = append(uniqf2lines, k)
				}
			}
			manageOutput(strings.Join(uniqf2lines, "\n"))
		}
	},
}

func init() {
	diffc.Flags().BoolVarP(&file1, "file1", "1", false, "Show Lines Unique to File1")
	diffc.Flags().BoolVarP(&file2, "file2", "2", false, "Show Lines Unique to File2")
	diffc.Flags().BoolVarP(&common, "common", "c", false, "Show Only Common Lines")

	diffc.CompletionOptions.DisableDefaultCmd = true
}
