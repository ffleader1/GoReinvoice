package direction

import "strings"

type Alignment string

type HorizontalAlignment Alignment
type VerticalAlignment Alignment

const HorizontalLeft HorizontalAlignment = "L"
const HorizontalCenter HorizontalAlignment = "C"
const HorizontalRight HorizontalAlignment = "R"

const VerticalTop VerticalAlignment = "T"
const VerticalMiddle VerticalAlignment = "M"
const VerticalBottom VerticalAlignment = "B"
const VerticalBaseline VerticalAlignment = "A"

func (al HorizontalAlignment) String() string {
	return string(al)
}

func (al VerticalAlignment) String() string {
	return string(al)
}

func ToHorizontalAlignment(str string) HorizontalAlignment {
	if IsHorizontalLeft(str) {
		return HorizontalLeft
	}

	if IsHorizontalRight(str) {
		return HorizontalRight
	}

	return HorizontalCenter
}

func ToVerticalAlignment(str string) VerticalAlignment {
	if IsVerticalTop(str) {
		return VerticalTop
	}

	if IsVerticalBottom(str) {
		return VerticalBottom
	}

	return VerticalMiddle
}

func IsVerticalTop(str string) bool {
	return strings.ToLower(strings.TrimSpace(str)) == "top"
}

func IsVerticalMiddle(str string) bool {
	return strings.ToLower(strings.TrimSpace(str)) == "middle"
}

func IsVerticalBottom(str string) bool {
	return strings.ToLower(strings.TrimSpace(str)) == "bottom"
}

func IsHorizontalLeft(str string) bool {
	return strings.ToLower(strings.TrimSpace(str)) == "left"
}

func IsHorizontalCenter(str string) bool {
	return strings.ToLower(strings.TrimSpace(str)) == "center"
}

func IsHorizontalRight(str string) bool {
	return strings.ToLower(strings.TrimSpace(str)) == "right"
}
