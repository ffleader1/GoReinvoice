package textgen

import (
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/direction"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/fpdfpoint"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/textconfig"
)

type TextObject struct {
	TopLeftCorner fpdfpoint.Point
	Width         float64
	Height        float64
	DisplayBorder bool
	textconfig.TextConfig
}

func GenerateTextObject(x, y, width, height float64, text string, fontSize, fontFamily int, hAlign, vAlign string, displayBorder bool) TextObject {
	return TextObject{
		TopLeftCorner: fpdfpoint.Point{X: x,
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

func (to TextObject) Translation(x, y float64) TextObject {
	to.TopLeftCorner = to.TopLeftCorner.Translation(x, y)
	return to
}
