package internal

import (
	"time"

	"golang.design/x/clipboard"
)

// ReadClipboard reads string data from clipboard
func ReadClipboard() string {
	data := clipboard.Read(clipboard.FmtText)

	return string(data)
}

// WriteClipboard writes given string data to clipboard
func WriteClipboard(data string) {
	go func() {
		<-clipboard.Write(clipboard.FmtText, []byte(data))
	}()

	time.Sleep(time.Duration(2) * time.Second)
}

// init clipboard package and check for availability
func init() {
	err := clipboard.Init()
	HandleError(err, "Failed to initialize clipboard")
}
