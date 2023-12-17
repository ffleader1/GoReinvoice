package cardinal

import "strings"

type Cardinal string

const Top Cardinal = "T"
const Bottom Cardinal = "B"
const Left Cardinal = "L"
const Right Cardinal = "R"

func (ca Cardinal) String() string {
	return string(ca)
}

func IsTop(str string) bool {
	return strings.Contains(str, Top.String())
}

func IsBottom(str string) bool {
	return strings.Contains(str, Bottom.String())
}

func IsLeft(str string) bool {
	return strings.Contains(str, Left.String())
}

func IsRight(str string) bool {
	return strings.Contains(str, Right.String())
}
