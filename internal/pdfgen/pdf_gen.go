package pdfgen

import (
	"GoReinvoice/internal/elementgen/tablegen"
	"GoReinvoice/internal/inputdata"
	"fmt"
	"github.com/go-pdf/fpdf"
	"log"
)

func GenPdf(input inputdata.PdfInput) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	fmt.Println("Ow vl")
	for _, e := range input.Elements {

		if e.Type == "table" {
			fmt.Println("Dang loop")
			tableData := input.Tables[e.ID]
			mergedCell, err := tablegen.GenerateTableLine(e.X, e.Y, e.Width, e.Height, tableData)
			if err != nil {
				log.Fatal(err)
				return
			}
			for _, v := range mergedCell {
				topLeft := v.TopLeftCorner()
				bottomRight := v.BottomRightCorner()
				w := bottomRight.X - topLeft.X
				h := bottomRight.Y - topLeft.Y
				pdf.SetXY(float64(topLeft.X), float64(topLeft.Y))
				pdf.CellFormat(float64(w), float64(h), "test", v.CardinalString(), 0,
					"CM", false, 0, "")
			}
		}
	}
	fmt.Println("Cbi write")
	if err := pdf.OutputFileAndClose("hello.pdf"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done write file")
}
