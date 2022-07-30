package pkg

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tarunKoyalwar/bninja/pkg/internal"
)

var (
	delimeter = ""              // delimeter to use
	fields    = ""              // feilds to display
	skip      = false           // skip lines without delimiter
	index     = sort.IntSlice{} // sorted indexes
)

// cut : vertical cut module
var cut = &cobra.Command{
	Use:   "cut",
	Short: "Similar to GNU cut but better",
	Long:  "Similar to GNU cut but with clipboard access,multicharacter delimeter and can cut multiple whitespaces",
	Example: `
// cut with delimeter as go.mod and use data from clipboard
bninja cut -d "go.mod" -f 2 -i
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		// ######### check data source ############
		input := ""
		var buff *bufio.Scanner

		if clipIn {
			input = internal.ReadClipboard()
		} else if len(args) > 0 {
			input = strings.Join(args, " ")
		} else {
			buff = asynchronousInput()
		}

		if input != "" {
			buff = bufio.NewScanner(strings.NewReader(input))
		}

		if buff == nil {
			internal.HandleError(fmt.Errorf("no data source"), "Input was not given")
		}

		// ############# check output flags ################
		var buffer bytes.Buffer
		var storeOutput bool
		if out != "" || clipOut {
			storeOutput = true
		}

		getindexes()

		if len(index) == 0 || delimeter == "" {
			return fmt.Errorf("fields or delimeter was not given of failed to extract")
		}

		// For each line in input
		for buff.Scan() {
			line := buff.Text()
			tempdata := strings.Split(line, delimeter)

			if len(tempdata) >= 2 {
				finalarr := []string{}

				for _, i := range index {
					if i-1 >= len(tempdata) {
						continue
					}
					finalarr = append(finalarr, tempdata[i-1])
				}
				outstring := strings.Join(finalarr, " ")
				if storeOutput {
					buffer.WriteString(outstring)
					if stdout {
						fmt.Println(outstring)
					}
				} else {
					fmt.Println(outstring)
				}
			} else {
				if !skip && len(tempdata) > 0 {
					outstring := tempdata[0] + "\n"
					if storeOutput {
						buffer.WriteString(outstring)
						if stdout {
							fmt.Print(outstring)
						}
					} else {
						fmt.Print(outstring)
					}
				}
			}
		}

		if storeOutput {
			if clipOut {
				internal.WriteClipboard(buffer.String())
			} else if out != "" {
				err := ioutil.WriteFile(out, buff.Bytes(), 0644)
				if err != nil {
					internal.HandleError(err, "Failed to write output to file %v", out)
				}
			}
		}

		return nil

	},
}

// getint : getInt From String
func getint(d string) int {
	n, err := strconv.Atoi(d)
	internal.HandleError(err, "Cannot get int from string", d)
	return n
}

// getrange : get range of feilds from string
func getrange(d string) []int {
	data := strings.Split(d, "-")
	head := getint(data[0])
	tail := getint(data[1])
	arr := []int{}
	for i := head; i <= tail; i++ {
		arr = append(arr, i)
	}
	return arr
}

func getindexes() {
	if strings.Contains(fields, ",") {
		data := strings.Split(fields, ",")
		for _, v := range data {
			if strings.Contains(v, "-") {
				index = append(index, getrange(v)...)
			} else {
				index = append(index, getint(v))
			}
		}
	} else {
		if strings.Contains(fields, "-") {
			index = append(index, getrange(fields)...)
		} else {
			index = []int{getint(fields)}
		}
	}
}

func init() {
	cut.Flags().StringVarP(&delimeter, "delim", "d", "", "Delimeter to Use(Supports Multiple Chars)")
	cut.Flags().StringVarP(&fields, "fields", "f", "", "fields/colums but it can be 1-3 or 1,5,6 or 1-3,8 (Inclusive)(starts from 1)")
	cut.Flags().BoolVar(&skip, "skip", false, "Skip Lines not containing delimeter")

	cut.MarkFlagRequired(delimeter)
	cut.MarkFlagRequired(fields)

	cut.CompletionOptions.DisableDefaultCmd = true

}
