package textgen

import (
	"GoReinvoice/internal/customtypes/direction"
	"GoReinvoice/internal/elementgen/resuable/pointgen"
)

type TextObject struct {
	TopLeftCorner       pointgen.Point
	Width               float64
	Height              float64
	Text                string
	DisplayBorder       bool
	HorizontalAlignment direction.HorizontalAlignment
	VerticalAlignment   direction.VerticalAlignment
}

func GenerateTextObject(x, y int, width float64, height float64, text, hAlign, vAlign string, displayBorder bool) TextObject {
	return TextObject{
		TopLeftCorner: pointgen.Point{X: x,
			Y: y},
		Width:               width,
		Height:              height,
		Text:                text,
		DisplayBorder:       displayBorder,
		HorizontalAlignment: direction.ToHorizontalAlignment(hAlign),
		VerticalAlignment:   direction.ToVerticalAlignment(vAlign),
	}
}

func (to TextObject) AlignmentString() string {
	return to.HorizontalAlignment.String() + to.VerticalAlignment.String()
}

func (to TextObject) WidthForFpdf() float64 {
	return to.Width
}

func (to TextObject) HeightForFpdf() float64 {
	line := 1
	for _, r := range to.Text {
		if r == '\n' {
			line++
		}
	}

	return to.Height / float64(line)
}

func (to TextObject) BorderString() string {
	if to.DisplayBorder {
		return "BTLR"
	}

	return ""
}
