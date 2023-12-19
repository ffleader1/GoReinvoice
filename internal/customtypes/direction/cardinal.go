package direction

import "strings"

type Cardinal string

const CardinalTop Cardinal = "T"
const CardinalBottom Cardinal = "B"
const CardinalLeft Cardinal = "L"
const CardinalRight Cardinal = "R"

func (ca Cardinal) String() string {
	return string(ca)
}

func IsCardinalTop(str string) bool {
	return strings.Contains(str, CardinalTop.String())
}

func IsCardinalBottom(str string) bool {
	return strings.Contains(str, CardinalBottom.String())
}

func IsCardinalLeft(str string) bool {
	return strings.Contains(str, CardinalLeft.String())
}

func IsCardinalRight(str string) bool {
	return strings.Contains(str, CardinalRight.String())
}
