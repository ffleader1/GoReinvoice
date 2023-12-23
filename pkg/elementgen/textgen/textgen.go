package textgen

import (
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/direction"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/textconfig"
	"github.com/ffleader1/GoReinvoice/pkg/elementgen/resuable/pointgen"
)

type TextObject struct {
	TopLeftCorner pointgen.Point
	Width         float64
	Height        float64
	DisplayBorder bool
	textconfig.TextConfig
}

func GenerateTextObject(x, y int, width float64, height float64, text string, fontSize, fontFamily int, hAlign, vAlign string, displayBorder bool) TextObject {
	return TextObject{
		TopLeftCorner: pointgen.Point{X: x,
			Y: y},
		Width:         width,
		Height:        height,
		DisplayBorder: displayBorder,
		TextConfig: textconfig.TextConfig{
			Text:                text,
			FontSize:            fontSize,
			FontFamily:          fontFamily,
			HorizontalAlignment: direction.ToHorizontalAlignment(hAlign),
			VerticalAlignment:   direction.ToVerticalAlignment(vAlign),
		},
	}
}

//func (to TextObject) AlignmentString() string {
//	return to.HorizontalAlignment.String() + to.VerticalAlignment.String()
//}

func (to TextObject) WidthForFpdf() float64 {
	return to.Width
}

func (to TextObject) HeightForFpdf() float64 { return to.Height }

func (to TextObject) BorderString() string {
	if to.DisplayBorder {
		return "BTLR"
	}

	return ""
}
