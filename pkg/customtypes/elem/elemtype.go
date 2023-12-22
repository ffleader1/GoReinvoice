package elem

import "strings"

type ElType string

const Table ElType = "table"
const Image ElType = "image"
const Text ElType = "text"
const Code128 ElType = "code128"
const Qrcode ElType = "qrcode"
const Line ElType = "line"

func (et ElType) String() string {
	return string(et)
}

func ToElemType(str string) ElType {
	str = strings.ToLower(strings.TrimSpace(str))
	switch str {
	case Table.String():
		return Table
	case Image.String():
		return Image
	case Text.String():
		return Text
	case Code128.String():
		return Code128
	case Qrcode.String():
		return Qrcode
	case Line.String():
		return Line
	}
	return ""
}
