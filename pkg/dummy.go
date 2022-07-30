package pkg

import "github.com/spf13/cobra"

/*
dummy :
Dummy module to pass data from input source to output source
*/
var dummy = &cobra.Command{
	Use:   "dummy",
	Short: "Dummy Module to pass data from input source to output ",
	Long:  "Dummy Moudle to pass data from input source(clipboard,file,stdin,args) to output(file,clipboard)",
	Example: `
// export stdin data to clipboard
echo "somestuff" | bninja dummy -o

//import data from clipboard
bninja dummy -i
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		input := synchronousInput(args)
		manageOutput(input)

		return nil
	},
}

func init() {
	dummy.CompletionOptions.DisableDefaultCmd = true
}
