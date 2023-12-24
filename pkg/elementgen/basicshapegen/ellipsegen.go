package basicshapegen

import "github.com/ffleader1/GoReinvoice/pkg/customtypes/fpdfpoint"

type EllipseObject struct {
	ID string
	fpdfpoint.Point
	RHorizontal float64
	RVertical   float64
	LineWidth   float64
	DegRotate   float64
}

func GenerateEllipseObject(id string, x, y, width, height, strokeWidth float64, defLineWidth float64, angle float64) EllipseObject {
	XFpdf := x + (width)/2
	YFpdf := y + (height)/2
	return EllipseObject{
		ID: id,
		Point: fpdfpoint.Point{
			X: XFpdf,
			Y: YFpdf},
		RHorizontal: width / 2,
		RVertical:   height / 2,
		LineWidth:   strokeWidth * defLineWidth,
		DegRotate:   angle,
	}
}

func (eo EllipseObject) Translation(x, y float64) EllipseObject {
	eo.Point = eo.Point.Translation(x, y)
	return eo
}
