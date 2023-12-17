package pdfgen

import (
	"GoReinvoice/internal/elementgen/tablegen"
	"GoReinvoice/internal/inputdata"
	"github.com/go-pdf/fpdf"
	"log"
)

func GenPdf(input inputdata.PdfInput) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	for _, e := range input.Elements {

		if e.Type == "table" {
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
		}
	}

	if err := pdf.OutputFileAndClose("hello.pdf"); err != nil {
		log.Fatal(err)
	}

}
