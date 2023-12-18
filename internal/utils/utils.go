package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

func GenerateFontStyleNum() map[int]string {
	styleMap := make(map[int]string)
	for i := 0; i < 1<<4; i++ {
		// Extract individual bits
		b0 := (i >> 3) & 1
		b1 := (i >> 2) & 1
		b2 := (i >> 1) & 1
		b3 := i & 1

		styleString := ""
		// Build the binary string
		binaryString := fmt.Sprintf("%d%d%d%d", b3, b2, b1, b0)

		if b3 == 1 {
			styleString += "S"
		}
		if b2 == 1 {
			styleString += "U"
		}
		if b1 == 1 {
			styleString += "I"
		}
		if b0 == 1 {
			styleString += "B"
		}

		// Convert to base 10
		decimalValue := 0
		for pos, bit := range binaryString {
			decimalValue += int(bit-'0') * (1 << (3 - pos))
		}

		// Print the binary and decimal representation
		styleMap[decimalValue] = styleString
	}
	return styleMap
}

func IsRelativePath(path string) bool {
	return !filepath.IsAbs(path) && !strings.HasPrefix(path, "http://") && !strings.HasPrefix(path, "https://") && !strings.HasPrefix(path, "data:")
}
