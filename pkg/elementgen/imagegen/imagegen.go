package imagegen

import (
	"bytes"
	"fmt"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/customerr"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/fpdfpoint"
	"github.com/go-pdf/fpdf"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type ImageObject struct {
	ID         string
	FpdfOption fpdf.ImageOptions
	Buffer     bytes.Buffer
	fpdfpoint.Point
	Width              float64
	DefaultScaleWidth  float64
	Height             float64
	DefaultScaleHeight float64
}

func GenerateImageObject(id string, imageUlr string, x, y, width, height float64, scale []float64) (ImageObject, error) {
	var option fpdf.ImageOptions

	file, err := os.Open(imageUlr)
	if err != nil {
		return ImageObject{}, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return ImageObject{}, err
	}

	var buf bytes.Buffer
	// Automatically choose encoder based on file extension
	ext := strings.ReplaceAll(filepath.Ext(imageUlr), ".", "")
	option.ImageType = ext
	switch ext {
	case "jpg", "jpeg":
		err = jpeg.Encode(&buf, img, nil)
	case "png":
		err = png.Encode(&buf, img)
	default:
		return ImageObject{}, fmt.Errorf("%w: %s", customerr.ErrUnsupportedImageFormat, ext)
	}
	if err != nil {
		return ImageObject{}, err
	}

	if len(scale) != 2 {
		return ImageObject{}, customerr.ErrInvalidScaleData
	}
	return ImageObject{
		ID:         id,
		FpdfOption: option,
		Buffer:     buf,
		Point: fpdfpoint.Point{
			X: x,
			Y: y,
		},
		Width:              width,
		DefaultScaleWidth:  scale[0],
		Height:             height,
		DefaultScaleHeight: scale[1],
	}, nil
}

func (io ImageObject) WidthForFpdf(scale ...float64) float64 {
	if scale == nil || len(scale) == 0 {
		return io.Width * io.DefaultScaleWidth
	}
	sc := 1.0
	for _, s := range scale {
		sc *= s
	}
	return io.Width * sc
}

func (io ImageObject) HeightForFpdf(scale ...float64) float64 {
	if scale == nil || len(scale) == 0 {
		return io.Height * io.DefaultScaleHeight
	}
	sc := 1.0
	for _, s := range scale {
		sc *= s
	}
	return io.Height * sc
}

func (io ImageObject) Translation(x, y float64) ImageObject {
	io.Point = io.Point.Translation(x, y)
	return io
}
