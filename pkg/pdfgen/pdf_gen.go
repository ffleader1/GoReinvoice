package pdfgen

import (
	"errors"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/elem"
	"github.com/ffleader1/GoReinvoice/pkg/elementgen/basicshapegen"
	"github.com/ffleader1/GoReinvoice/pkg/elementgen/codegen"
	"github.com/ffleader1/GoReinvoice/pkg/elementgen/imagegen"
	"github.com/ffleader1/GoReinvoice/pkg/elementgen/tablegen"
	"github.com/ffleader1/GoReinvoice/pkg/elementgen/textgen"
	"github.com/ffleader1/GoReinvoice/pkg/inputdata"
	"github.com/ffleader1/GoReinvoice/pkg/utils"
	"github.com/go-pdf/fpdf"
	"log"
	"strings"
)

type fontConfig struct {
	familyName string
	style      string
}

const arialFontPrefix = 1000
const timesFontPrefix = 2000

type PdfData struct {
	pdf                   *fpdf.Fpdf
	pdfData               inputdata.PdfInput
	pdfDefaultStrokeWidth float64
	fonts                 map[int]fontConfig
}

func NewPdfDataFromFile(file string) (PdfData, error) {
	input, err := inputdata.ReadData(file)
	if err != nil {
		return PdfData{}, err
	}
	pdfData := NewPdfData(input)
	return pdfData, nil
}

func NewPdfData(pdfData inputdata.PdfInput) PdfData {
	pdf := fpdf.New(pdfData.Paper.Orientation, pdfData.Paper.Unit, pdfData.Paper.Size, "")
	var fontMap = make(map[int]fontConfig)
	const (
		defaultFont  = "Arial"
		defaultStyle = "B"
		defaultSize  = 16.0
	)
	pdf.SetFont(defaultFont, defaultStyle, defaultSize)
	fontMap[0] = fontConfig{
		familyName: defaultFont,
		style:      defaultStyle,
	}
	for k, v := range utils.GenerateFontStyleNum() {
		fontMap[arialFontPrefix+k] = fontConfig{
			familyName: "Arial",
			style:      v,
		}

		fontMap[timesFontPrefix+k] = fontConfig{
			familyName: "Times",
			style:      v,
		}
	}
	for index, f := range pdfData.Fonts {
		pdf.AddUTF8Font(f.FamilyName, f.Style, f.DataURL)
		fontMap[index+1] = fontConfig{
			familyName: f.FamilyName,
			style:      f.Style,
		}
	}
	pdf.AddPage()

	return PdfData{
		pdf:                   pdf,
		pdfData:               pdfData,
		pdfDefaultStrokeWidth: pdf.GetLineWidth(),
		fonts:                 fontMap,
	}
}

func (pd *PdfData) SwitchFont(fontStyleIndex int) (string, string) {
	fd, ok := pd.fonts[fontStyleIndex]
	if ok {
		return fd.familyName, fd.style
	}
	return "", ""
}

func (pd *PdfData) GenPdf(placeHolderMap map[string]string, outputFile string) {

	//pdf.AddUTF8Font()
	pdf := *pd.pdf
	for _, e := range pd.pdfData.Elements {

		switch elem.ToElemType(e.Type) {
		case elem.Table:

			tableData := pd.pdfData.Tables[e.ID]
			mergedCell, err := tablegen.GenerateCellMap(e.X, e.Y, int(e.Width), int(e.Height), e.StrokeWidth, tableData)
			if err != nil {
				log.Fatal(err)
				return
			}
			for _, cell := range mergedCell {
				topLeft := cell.TopLeftCorner()
				pdf.SetXY(float64(topLeft.X), float64(topLeft.Y))
				pdf.MultiCell(cell.WidthForFpdf(), cell.HeightForFpdf(), pd.fillPlaceHolder(cell.Text, placeHolderMap), cell.CardinalString(),
					"CM", false)
			}

		case elem.Image:
			var file inputdata.File
			file, found := pd.pdfData.Files[e.ID]
			if !found {
				log.Println(errors.New("file config not found"))
				continue
			}
			imageObject, err := imagegen.GenerateImageObject(file.DataURL, e.X, e.Y, e.Width, e.Height, e.Scale)
			if err != nil {
				log.Println(err)
				continue
			}
			pdf.RegisterImageOptionsReader(imageObject.Name, imageObject.FpdfOption, &imageObject.Buffer)
			pdf.Image(imageObject.Name, float64(imageObject.X), float64(imageObject.Y), imageObject.WidthForFpdf(), imageObject.HeightForFpdf(), false, "", 0, "")
		case elem.Text:
			textObject := textgen.GenerateTextObject(e.X, e.Y, e.Width, e.Height, pd.fillPlaceHolder(e.Text, placeHolderMap), e.TextAlign, e.VerticalAlign, false)

			font, style := pd.SwitchFont(e.FontFamily)
			pdf.SetFont(font, style, float64(e.FontSize))

			pdf.SetXY(float64(textObject.TopLeftCorner.X), float64(textObject.TopLeftCorner.Y))

			pdf.MultiCell(textObject.WidthForFpdf(), textObject.HeightForFpdf(), textObject.Text, textObject.BorderString(),
				textObject.AlignmentString(), false)

		case elem.Code128, elem.Qrcode:
			codeObject, err := codegen.GenerateCodeObject(e.Type, pd.fillPlaceHolder(e.Text, placeHolderMap), e.X, e.Y)
			if err != nil {
				continue
			}
			pdf.RegisterImageOptionsReader(codeObject.Name, codeObject.FpdfOption, &codeObject.Buffer)
			pdf.Image(codeObject.Name, float64(e.X), float64(e.Y), e.Width, e.Height, false, "", 0, "")
		case elem.Line:
			lineObject, err := basicshapegen.GenerateLineObject(e.X, e.Y, e.Point, e.StrokeWidth, pd.pdfDefaultStrokeWidth)
			if err != nil {
				continue
			}
			pdf.SetLineWidth(lineObject.LineWidth)
			pdf.Line(lineObject.StartX, lineObject.StartY, lineObject.EndX, lineObject.EndY)
		case elem.Ellipse:
			ellipseObject := basicshapegen.GenerateEllipseObject(e.X, e.Y, e.Width, e.Height, e.StrokeWidth, pd.pdfDefaultStrokeWidth, e.Angle)
			pdf.SetLineWidth(ellipseObject.LineWidth)
			pdf.Ellipse(ellipseObject.X, ellipseObject.Y, ellipseObject.RHorizontal, ellipseObject.RVertical, ellipseObject.DegRotate, "D")
		}
	}

	if err := pdf.OutputFileAndClose(outputFile); err != nil {
		log.Fatal(err)
	}
}

func (pd *PdfData) fillPlaceHolder(toFill string, placeHolderMap map[string]string) string {
	for k, v := range placeHolderMap {
		holder := "{{" + k + "}}"
		if strings.Contains(toFill, holder) {
			toFill = strings.ReplaceAll(toFill, holder, v)
		}
	}

	return toFill
}
