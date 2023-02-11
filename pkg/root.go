package pkg

import (
	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:               "bninja",
	Short:             "A Handy Toolkit for bash ninja's",
	Long:              "Toolkit containing Implementation of popular unix commands with extra features to imporve efficiency of shell users ",
	DisableAutoGenTag: true,
}

func Execute() {
	err := rootCMD.Execute()
	if err != nil {
		panic(err)
	}
}

func init() {

	rootCMD.AddCommand(diffc)
	rootCMD.AddCommand(dummy)
	rootCMD.AddCommand(getlenc)
	rootCMD.AddCommand(hcut)
	rootCMD.AddCommand(htmlenc)
	rootCMD.AddCommand(jsonn)
	rootCMD.AddCommand(replace)
	rootCMD.AddCommand(sortc)
	rootCMD.AddCommand(unique)
	rootCMD.AddCommand(urlEnc)
	rootCMD.AddCommand(cut)
	rootCMD.AddCommand(Filter)

	rootCMD.PersistentFlags().BoolVarP(&clipIn, "clipin", "i", false, "Import data from clipboard")
	rootCMD.PersistentFlags().BoolVarP(&clipOut, "clipout", "o", false, "Export data to clipboard")
	rootCMD.PersistentFlags().StringVar(&file, "file", "", "Import data from File")
	rootCMD.PersistentFlags().BoolVar(&stdout, "stdout", false, "Print to Stdout even if data if data is exported to clipboard or file")
	rootCMD.PersistentFlags().StringVar(&out, "out", "", "Export data to File")

	rootCMD.CompletionOptions.DisableDefaultCmd = true
}
