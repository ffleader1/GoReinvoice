package codegen

import (
	"bytes"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/elem"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/fpdfpoint"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/textconfig"
	"github.com/go-pdf/fpdf"
	"image/png"
)

type CodeObject struct {
	ID         string
	Content    string
	FpdfOption fpdf.ImageOptions
	Buffer     bytes.Buffer
	fpdfpoint.Point
	Width  float64
	Height float64
}

func GenerateCodeObject(id string, codeType string, data string, x, y, w, h float64, placeHolderMap map[string]string) (CodeObject, error) {
	var qrCode barcode.Barcode
	var err error
	textCfg := textconfig.TextConfig{
		Text:                data,
		FontSize:            0,
		FontFamily:          0,
		HorizontalAlignment: "",
		VerticalAlignment:   "",
	}
	replacedData := textCfg.TextWithPlaceholder(placeHolderMap)
	switch elem.ToElemType(codeType) {
	case elem.Code128:
		if qrCode, err = code128.Encode(replacedData); err != nil {
			return CodeObject{}, err
		}
	case elem.Qrcode:
		if qrCode, err = qr.Encode(replacedData, qr.M, qr.Auto); err != nil {
			return CodeObject{}, err
		}
		if qrCode, err = barcode.Scale(qrCode, 768, 768); err != nil {
			return CodeObject{}, err
		}
	}

	var buffer bytes.Buffer

	if err = png.Encode(&buffer, qrCode); err != nil {
		return CodeObject{}, err
	}

	return CodeObject{ID: id,
		Content: replacedData,
		FpdfOption: fpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: "png",
		},
		Buffer: buffer,
		Point: fpdfpoint.Point{
			X: x,
			Y: y,
		},
		Width:  w,
		Height: h,
	}, nil
}

func (co CodeObject) Translation(x, y float64) CodeObject {
	co.Point = co.Point.Translation(x, y)
	return co
}
