package pdfgen

import (
	"GoReinvoice/internal/elementgen/tablegen"
	"GoReinvoice/internal/inputdata"
	"GoReinvoice/internal/utils"
	"github.com/go-pdf/fpdf"
	"log"
)

type fontConfig struct {
	familyName string
	style      string
}

const arialFontPrefix = 1000
const timesFontPrefix = 2000

type PdfData struct {
	pdf   *fpdf.Fpdf
	fonts map[int]fontConfig
}

func (pd *PdfData) NewPdfData(pdfData inputdata.PdfInput) PdfData {
	pdf := fpdf.New(pdfData.Paper.Orientation, pdfData.Paper.Unit, pdfData.Paper.Size, "")
	var fontMap = make(map[int]fontConfig)
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
		fontMap[index] = fontConfig{
			familyName: f.FamilyName,
			style:      f.Style,
		}
	}
	pdf.SetFont("Arial", "B", 16)
	return PdfData{
		pdf:   pdf,
		fonts: fontMap,
	}
}

func (pd *PdfData) SwitchFont(fontStyleIndex int, size float64) {
	fd, ok := pd.fonts[fontStyleIndex]
	if ok {
		pd.pdf.SetFont(fd.familyName, fd.style, size)
	}
}

func GenPdf(input inputdata.PdfInput) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	//pdf.AddUTF8Font()

	for _, e := range input.Elements {

		switch e.Type {
		case "table":

			tableData := input.Tables[e.ID]
			mergedCell, err := tablegen.GenerateCellMap(e.X, e.Y, e.Width, e.Height, e.StrokeWidth, tableData)
			if err != nil {
				log.Fatal(err)
				return
			}
			for _, cell := range mergedCell {
				topLeft := cell.TopLeftCorner()
				pdf.SetXY(float64(topLeft.X), float64(topLeft.Y))
				pdf.MultiCell(cell.WidthForFpdf(), cell.HeightForFpdf(), cell.Text, cell.CardinalString(),
					"CM", false)
			}

		case "image":
			file := input.Files[e.ID]
			pdf.Image(file.DataURL, float64(e.X), float64(e.Y), float64(e.Width), float64(e.Height), false, "", 0, "")
		case "text":

		}

	}

	if err := pdf.OutputFileAndClose("hello.pdf"); err != nil {
		log.Fatal(err)
	}

}
