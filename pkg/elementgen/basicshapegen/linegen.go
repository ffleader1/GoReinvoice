package basicshapegen

import "errors"

type LineObject struct {
	StartX    float64
	StartY    float64
	EndX      float64
	EndY      float64
	LineWidth float64
}

func GenerateLineObject(x, y int, point [][]int, strokeWidth int, defLineWidth float64) (LineObject, error) {
	if len(point) != 2 {
		return LineObject{}, errors.New("invalid points length")
	}
	startX := float64(x + point[0][0])
	startY := float64(y + point[0][1])
	endX := float64(x + point[1][0])
	endY := float64(y + point[1][1])

	return LineObject{
		StartX:    startX,
		StartY:    startY,
		EndX:      endX,
		EndY:      endY,
		LineWidth: float64(strokeWidth) * defLineWidth,
	}, nil
}
