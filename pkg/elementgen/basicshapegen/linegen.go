package basicshapegen

import (
	"errors"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/fpdfpoint"
)

type LineObject struct {
	fpdfpoint.Line
}

func GenerateLineObject(x, y float64, point [][]float64, strokeWidth float64, defLineWidth float64) (LineObject, error) {
	if len(point) != 2 {
		return LineObject{}, errors.New("invalid points length")
	}
	startX := x + point[0][0]
	startY := y + point[0][1]
	endX := x + point[1][0]
	endY := y + point[1][1]

	return LineObject{
		Line: fpdfpoint.Line{
			A: fpdfpoint.Point{
				X: startX,
				Y: startY,
			},
			B: fpdfpoint.Point{
				X: endX,
				Y: endY,
			},
			Width:  strokeWidth * defLineWidth,
			Status: fpdfpoint.Visible,
		},
	}, nil
}

func (lo LineObject) Translation(x, y float64) LineObject {
	return LineObject{lo.Line.Translation(x, y)}
}
