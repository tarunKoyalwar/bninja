package pkg

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tarunKoyalwar/bninja/pkg/internal"
)

// Global Flags

var (
	clipIn  = false // clipin imports data from clipboard
	clipOut = false // clipout exports data from clipboard
	file    = ""    // read data from file with filename
	stdout  = false // must print on stdout
	out     = ""    // Save output/results to file with filename
)

// Managing I/O of all commands

// synchronousInput : Reads all data at once
func synchronousInput(args []string) string {

	if clipIn {
		return internal.ReadClipboard()
	} else if internal.HasStdin() {
		return internal.ReadAllStdin()
	} else if len(args) > 0 {
		return strings.Join(args, " ")
	} else if file != "" {
		return internal.ReadFromFile(file)
	} else {
		internal.HandleError(fmt.Errorf("no data source"), "Input was not given")
	}

	return ""
}

// asynchronousInput : read data using bufio.scanner
func asynchronousInput() *bufio.Scanner {
	var scanner *bufio.Scanner

	if internal.HasStdin() {
		scanner = bufio.NewScanner(os.Stdin)

	} else if file != "" {
		f, err := os.Open(file)
		internal.HandleError(err, "failed to read input from %v", file)
		scanner = bufio.NewScanner(f)
	}

	return scanner
}

// manageOutput : Write output based on given flags
func manageOutput(data string) {
	if clipOut {
		internal.WriteClipboard(data)
		if stdout {
			fmt.Println(data)
		}
	} else if out != "" {
		if stdout {
			fmt.Println(data)
		}
		err := ioutil.WriteFile(out, []byte(data), 0644)
		if err != nil {
			internal.HandleError(err, "Failed to write output to file %v", out)
		}
	} else {
		fmt.Println(data)
	}
}
