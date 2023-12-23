package basicshapegen

type EllipseObject struct {
	X           float64
	Y           float64
	RHorizontal float64
	RVertical   float64
	LineWidth   float64
	DegRotate   float64
}

func GenerateEllipseObject(x, y int, width, height float64, strokeWidth int, defLineWidth float64, angle float64) EllipseObject {
	XFpdf := float64(x) + (width)/2
	YFpdf := float64(y) + (height)/2
	return EllipseObject{
		X:           XFpdf,
		Y:           YFpdf,
		RHorizontal: width,
		RVertical:   height,
		LineWidth:   float64(strokeWidth) * defLineWidth,
		DegRotate:   angle,
	}
}
