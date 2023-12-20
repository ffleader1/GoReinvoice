package pdfgen

import (
	"GoReinvoice/internal/elementgen/codegen"
	"GoReinvoice/internal/elementgen/tablegen"
	"GoReinvoice/internal/elementgen/textgen"
	"GoReinvoice/internal/inputdata"
	"GoReinvoice/internal/utils"
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

		switch e.Type {
		case "table":

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

		case "image":
			file := pd.pdfData.Files[e.ID]
			pdf.Image(file.DataURL, float64(e.X), float64(e.Y), e.Width, e.Height, false, "", 0, "")
		case "text":
			textObject := textgen.GenerateTextObject(e.X, e.Y, e.Width, e.Height, pd.fillPlaceHolder(e.Text, placeHolderMap), e.TextAlign, e.VerticalAlign, false)

			font, style := pd.SwitchFont(e.FontFamily)
			pdf.SetFont(font, style, float64(e.FontSize))

			pdf.SetXY(float64(textObject.TopLeftCorner.X), float64(textObject.TopLeftCorner.Y))

			pdf.MultiCell(textObject.WidthForFpdf(), textObject.HeightForFpdf(), textObject.Text, textObject.BorderString(),
				textObject.AlignmentString(), false)

		case "code128", "qrcode":
			codeObject, err := codegen.GenerateCodeObject(e.Type, pd.fillPlaceHolder(e.Text, placeHolderMap))
			if err != nil {
				continue
			}
			options := fpdf.ImageOptions{
				ReadDpi:   false,
				ImageType: "png",
			}
			randName := utils.RandStringBytes(6)
			pdf.RegisterImageOptionsReader(randName, options, &codeObject.Buffer)
			pdf.Image(randName, float64(e.X), float64(e.Y), e.Width, e.Height, false, "", 0, "")
		case "line":
			startX := e.X + e.Point[0][0]
			startY := e.Y + e.Point[0][1]
			endX := e.X + e.Point[1][0]
			endY := e.Y + e.Point[1][1]
			pdf.SetLineWidth(pd.pdfDefaultStrokeWidth * float64(e.StrokeWidth))
			pdf.Line(float64(startX), float64(startY), float64(endX), float64(endY))
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
