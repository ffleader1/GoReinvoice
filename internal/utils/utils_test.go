package utils

import (
	"fmt"
	"testing"
)

func TestParseMarkdown(t *testing.T) {
	input := "This is __not normal__ text. **This is bold text.**__This is italic text.__ Hello, world||123||!"
	segments := ParseMarkdown(input)

	for _, segment := range segments {
		fmt.Printf("Text: `%s`, Style: %s\n", segment.Text, segment.Style)
	}
}
