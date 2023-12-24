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
	objectData            DataObject
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
		objectData:            NewDataObject(),
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
		pd.GenElement(&pdf, elem.ToElemType(e.Type), false, &e, nil, placeHolderMap)
	}

	if err := pdf.OutputFileAndClose(outputFile); err != nil {
		log.Fatal(err)
	}
}
func (pd *PdfData) GenElement(pdf *fpdf.Fpdf, elType elem.ElType, copying bool, e *inputdata.Element, data interface{}, placeHolderMap map[string]string) {
	var err error
	switch elType {
	case elem.Table:
		var tableObject tablegen.TableObject
		if !copying {
			tableData, found := pd.pdfData.Tables[e.ID]
			if !found {
				log.Println(errors.New("table data config not found"))
				return
			}
			tableObject, err = tablegen.GenerateTableObject(e.ID, e.X, e.Y, e.Width, e.Height, e.StrokeWidth, tableData)
			if err != nil {
				log.Println(err)
				return
			}
			pd.objectData.Save(e.ID, elType, tableObject)
		} else {
			tableObject = data.(tablegen.TableObject).Translation(e.X, e.Y)
		}
		for _, cell := range tableObject.CellMap {
			topLeft := cell.TopLeftCorner()
			font, style := pd.SwitchFont(cell.FontFamily)
			pdf.SetFont(font, style, float64(cell.FontSize))

			pdf.SetXY(topLeft.X, topLeft.Y)
			pdf.MultiCell(cell.WidthForFpdf(), cell.HeightForFpdf(), cell.TextWithPlaceholder(placeHolderMap), cell.CardinalString(),
				cell.AlignmentString(), false)
		}

	case elem.Image:
		var imageObject imagegen.ImageObject
		if !copying {
			file, found := pd.pdfData.Files[e.ID]
			if !found {
				log.Println(errors.New("file config not found"))
				return
			}
			imageObject, err = imagegen.GenerateImageObject(e.ID, file.DataURL, e.X, e.Y, e.Width, e.Height, e.Scale)
			if err != nil {
				log.Println(err)
				return
			}
			pdf.RegisterImageOptionsReader(imageObject.ID, imageObject.FpdfOption, &imageObject.Buffer)
			pd.objectData.Save(e.ID, elType, imageObject)
		} else {
			imageObject = data.(imagegen.ImageObject).Translation(e.X, e.Y)
		}

		pdf.Image(imageObject.ID, imageObject.X, imageObject.Y, imageObject.WidthForFpdf(), imageObject.HeightForFpdf(), false, "", 0, "")

	case elem.Text:
		var textObject textgen.TextObject
		if !copying {
			textObject = textgen.GenerateTextObject(e.ID, e.X, e.Y, e.Width, e.Height, e.Text, e.FontSize, e.FontFamily, e.TextAlign, e.VerticalAlign, false)
			pd.objectData.Save(e.ID, elType, textObject)
		} else {
			textObject = data.(textgen.TextObject).Translation(e.X, e.Y)
		}
		font, style := pd.SwitchFont(textObject.FontFamily)
		pdf.SetFont(font, style, float64(textObject.FontSize))

		pdf.SetXY(textObject.TopLeftCorner.X, textObject.TopLeftCorner.Y)

		pdf.MultiCell(textObject.WidthForFpdf(), textObject.HeightForFpdf(), textObject.TextWithPlaceholder(placeHolderMap), textObject.BorderString(),
			textObject.AlignmentString(), false)

	case elem.Code128, elem.Qrcode:
		var codeObject codegen.CodeObject
		if !copying {
			codeObject, err = codegen.GenerateCodeObject(e.ID, e.Type, e.Text, e.X, e.Y, e.Width, e.Height, placeHolderMap)
			if err != nil {
				log.Println(err)
				return
			}
			pdf.RegisterImageOptionsReader(codeObject.ID, codeObject.FpdfOption, &codeObject.Buffer)
			pd.objectData.Save(e.ID, elType, codeObject)
		} else {
			codeObject = data.(codegen.CodeObject).Translation(e.X, e.Y)
		}
		pdf.Image(codeObject.ID, codeObject.X, codeObject.Y, codeObject.Width, codeObject.Height, false, "", 0, "")
	case elem.Line:
		var lineObject basicshapegen.LineObject
		if !copying {
			lineObject, err = basicshapegen.GenerateLineObject(e.ID, e.X, e.Y, e.Points, e.StrokeWidth, pd.pdfDefaultStrokeWidth)
			if err != nil {
				log.Println(err)
				return
			}
			pd.objectData.Save(e.ID, elType, lineObject)
		} else {
			lineObject = data.(basicshapegen.LineObject).Translation(e.X, e.Y)
		}

		pdf.SetLineWidth(lineObject.Width)
		pdf.Line(lineObject.A.X, lineObject.A.Y, lineObject.B.X, lineObject.B.Y)
	case elem.Ellipse:
		var ellipseObject basicshapegen.EllipseObject
		if !copying {
			ellipseObject = basicshapegen.GenerateEllipseObject(e.ID, e.X, e.Y, e.Width, e.Height, e.StrokeWidth, pd.pdfDefaultStrokeWidth, e.Angle)
			pd.objectData.Save(e.ID, elType, ellipseObject)
		} else {
			ellipseObject = data.(basicshapegen.EllipseObject).Translation(e.X, e.Y)
		}
		pdf.SetLineWidth(ellipseObject.LineWidth)
		pdf.Ellipse(ellipseObject.X, ellipseObject.Y, ellipseObject.RHorizontal, ellipseObject.RVertical, ellipseObject.DegRotate, "D")
	case elem.Copy:
		if copying {
			return
		}
		copyData, found := pd.pdfData.Copies[e.ID]
		if !found {
			log.Println(errors.New("table data config not found"))
			return
		}
		for _, c := range copyData {
			if tp, dat, err := pd.objectData.Load(c); err == nil {
				pd.GenElement(pdf, tp, true, e, dat, placeHolderMap)
			} else {
				log.Println(err)
			}
		}
	}

}
