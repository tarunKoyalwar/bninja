package internal

import (
	"fmt"
	"io"
	"log"
	"os"
)

// ReadAllStdin : Read data from stdin
func ReadAllStdin() string {
	bin, err := io.ReadAll(os.Stdin)
	HandleError(err, "failed to read data from stdin")
	return string(bin)
}

// ReadFromFile : Reads data from file
func ReadFromFile(filename string) string {
	bin, err := os.ReadFile(filename)
	HandleError(err, "failed to read data from file `%v`", filename)
	return string(bin)
}

func HandleError(er error, format string, a ...any) {
	if er != nil {
		fmt.Printf(format+"\n", a...)
		log.Fatal(er)
	}
}

// HasStdin : Check if Stdin is present
func HasStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	mode := stat.Mode()

	isPipedFromChrDev := (mode & os.ModeCharDevice) == 0
	isPipedFromFIFO := (mode & os.ModeNamedPipe) != 0

	return isPipedFromChrDev || isPipedFromFIFO
}
