package utils

import (
	"fmt"
	"slices"
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

type TextStyle struct {
	Text  string
	Style string // 0: Normal, 1: Bold, 2: Italic
}

func ParseMarkdown(text string) []TextStyle {

	newStr, idxBold := replaceWithIndex(text, "**", "~~")
	newStr2, idxItalic := replaceWithIndex(newStr, "__", "~~")
	newStr3, idxItalicAndBold := replaceWithIndex(newStr2, "||", "~~")
	newStr4, splitIdx := splitWithIndexes(newStr3, "~~")

	ts := make([]TextStyle, 0)
	for idx, n := range newStr4 {
		if slices.Contains(idxBold, splitIdx[idx]) {
			ts = append(ts, TextStyle{n, "B"})
			continue
		}
		if slices.Contains(idxItalic, splitIdx[idx]) {
			ts = append(ts, TextStyle{n, "I"})
			continue
		}
		if slices.Contains(idxItalicAndBold, splitIdx[idx]) {
			ts = append(ts, TextStyle{n, "BI"})
			continue
		}
		ts = append(ts, TextStyle{n, ""})
	}

	return ts

}
func splitWithIndexes(str string, sep string) ([]string, []int) {
	var parts []string
	var indexes []int
	startIndex := 0
	for i := 0; i <= len(str)-len(sep); i++ {
		if strings.Compare(str[i:i+len(sep)], sep) == 0 {
			if startIndex < i {
				parts = append(parts, str[startIndex:i])
				indexes = append(indexes, startIndex)
			}
			startIndex = i + len(sep)
		}
	}
	if startIndex < len(str) {
		parts = append(parts, str[startIndex:])
		indexes = append(indexes, startIndex)
	}
	return parts, indexes
}

func replaceWithIndex(s string, old, new string) (string, []int) {
	var result string
	var indexes []int
	offset := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if strings.HasPrefix(s[i:], old) {
			if count%2 == 0 {
				indexes = append(indexes, i+len(new))
			}
			count++
			result += new
			i += len(new) - 1
			offset += len(old)
		} else {
			result += string(s[i])
		}
	}
	return result, indexes
}
