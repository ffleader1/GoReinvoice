package pdfdata

import "github.com/ffleader1/GoReinvoice/pkg/pdfgen"

type PDFData struct {
	*pdfgen.PdfData
}

func NewPDFDataFromFile(file string) (PDFData, error) {
	data, err := pdfgen.NewPdfDataFromFile(file)
	if err != nil {
		return PDFData{nil}, err
	}
	return PDFData{&data}, nil
}

func (pd *PDFData) GenPDFToFile(placeHolderMap map[string]string, outputFile string) {
	pd.PdfData.GenPdf(placeHolderMap, outputFile)
}
