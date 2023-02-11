package pkg

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tarunKoyalwar/bninja/pkg/internal"
)

var (
	KeyWords    = []string{}
	Insensitive = false
	wordFile    = ""
)

var Filter = &cobra.Command{
	Use:   "filter",
	Short: "filters data from input using keywords",
	Long:  "filters out lines that contains any of the given keyword.Note filter does not support clipboard input",
	RunE: func(cmd *cobra.Command, args []string) error {
		scanner := asynchronousInput()
		if scanner == nil {
			log.Fatal("no input provided")
		}
		data := internal.ReadFromFile(wordFile)
		KeyWords = append(KeyWords, strings.Split(data, "\n")...)
		if len(KeyWords) == 0 {
			log.Fatal("filter keywords not provided")
		}

		for scanner.Scan() {
			line := scanner.Text()
			if !FilterLine(line) {
				fmt.Println(line)
			}
		}
		return nil
	},
}

// FilterLine Uses line
func FilterLine(line string) bool {
	if line == "" {
		return true
	}
	if Insensitive {
		line = strings.ToLower(line)
	}
	for _, v := range KeyWords {
		if strings.Contains(line, v) {
			return true
		}
	}
	return false
}

func init() {
	Filter.Flags().StringVarP(&wordFile, "keywords", "k", "", "Keywords File contains keywords to filter")
	Filter.MarkFlagRequired("keywords")
	Filter.CompletionOptions.DisableDefaultCmd = true
}
