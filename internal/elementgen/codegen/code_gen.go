package codegen

import (
	"bytes"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"strings"
)

type CodeObject struct {
	Text   string
	Buffer bytes.Buffer
}

func GenerateCodeObject(codeType string, data string) (CodeObject, error) {
	var qrCode barcode.Barcode
	var err error
	switch strings.ToLower(strings.TrimSpace(codeType)) {
	case "code128":
		if qrCode, err = code128.Encode(data); err != nil {
			return CodeObject{}, err
		}
	case "qrcode":
		if qrCode, err = qr.Encode(data, qr.M, qr.Auto); err != nil {
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

	return CodeObject{Text: data, Buffer: buffer}, nil
}
