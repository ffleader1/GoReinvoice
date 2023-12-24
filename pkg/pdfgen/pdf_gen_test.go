package pdfgen

import (
	"github.com/ffleader1/GoReinvoice/pkg/inputdata"
	"log"
	"testing"
)

func TestGenTablePdf(t *testing.T) {
	input := inputdata.PdfInput{
		Type:       "test",
		Version:    0,
		Source:     "",
		SourceSize: "",
		Rotation:   "",
		Elements: []inputdata.Element{
			{
				ID:          "elm1",
				Type:        "table",
				X:           20,
				Y:           30,
				Width:       180,
				Height:      100,
				StrokeWidth: 5,
			},
		},
		AppState: inputdata.AppState{},
		Files:    nil,
		Tables: map[string]inputdata.Table{
			"elm1": {
				ColumnRatio: []float64{0.2, 0.2, 0.2, 0.2, 0.2},
				RowRatio:    []float64{0.25, 0.25, 0.25, 0.25},
				MergeCell:   map[string]string{"A3": "C4", "D1": "E1"},
				HiddenEdge:  map[string]string{"A3": "LB", "E4": "TR"},
				CellText:    map[string]inputdata.TextConfig{"A3": {Text: "O nay\nsiu to\n...", FontSize: 16, FontFamily: 0}, "B1": {Text: "o nay be", FontSize: 16, FontFamily: 0}},
			},
		},
		Paper: inputdata.Paper{
			Size:        "A4",
			Unit:        "mm",
			Orientation: "P",
		},
	}
	pdfData := NewPdfData(input)
	pdfData.GenPdf(nil, "test_gen_table.pdf")
}

func TestGenImagePdf(t *testing.T) {
	input := inputdata.PdfInput{
		Type:       "test",
		Version:    0,
		Source:     "",
		SourceSize: "",
		Rotation:   "",
		Elements: []inputdata.Element{
			{
				ID:     "elm1",
				Type:   "image",
				X:      20,
				Y:      30,
				Width:  120,
				Height: 80,
			},
		},
		AppState: inputdata.AppState{},
		Files: map[string]inputdata.File{
			"elm1": {
				MimeType: "image/png",
				DataURL:  "../../resource/image/go_rust.png",
			},
		},
		Paper: inputdata.Paper{
			Size:        "A4",
			Unit:        "mm",
			Orientation: "P",
		},
	}
	pdfData := NewPdfData(input)
	pdfData.GenPdf(nil, "test_gen_image.pdf")
}

func TestGenStringPdf(t *testing.T) {
	input := inputdata.PdfInput{
		Type:       "test",
		Version:    0,
		Source:     "",
		SourceSize: "",
		Rotation:   "",
		Elements: []inputdata.Element{
			{
				ID:     "elm1",
				Type:   "text",
				X:      20,
				Y:      30,
				Width:  120,
				Height: 60,
				TextConfig: inputdata.TextConfig{FontSize: 20,
					FontFamily: 1001,
					Text:       "Oc cho is real",
				},
			},
		},
		AppState: inputdata.AppState{},
		Files:    nil,
		Tables:   nil,
		Paper: inputdata.Paper{
			Size:        "A4",
			Unit:        "mm",
			Orientation: "P",
		},
	}
	pdfData := NewPdfData(input)
	pdfData.GenPdf(nil, "test_gen_text.pdf")
}

func TestGenCodePdf(t *testing.T) {
	input := inputdata.PdfInput{
		Type:       "test",
		Version:    0,
		Source:     "",
		SourceSize: "",
		Rotation:   "",
		Elements: []inputdata.Element{
			{
				ID:         "elm1",
				Type:       "code128",
				X:          20,
				Y:          30,
				Width:      120,
				Height:     60,
				TextConfig: inputdata.TextConfig{Text: "Testing Code 128 Generation"},
			},
		},
		AppState: inputdata.AppState{},
		Files:    nil,
		Tables:   nil,
		Paper: inputdata.Paper{
			Size:        "A4",
			Unit:        "mm",
			Orientation: "P",
		},
	}
	pdfData := NewPdfData(input)
	pdfData.GenPdf(nil, "test_gen_code.pdf")
}

func TestGenShapePdf(t *testing.T) {
	input := inputdata.PdfInput{
		Type:       "test",
		Version:    0,
		Source:     "",
		SourceSize: "",
		Rotation:   "",
		Elements: []inputdata.Element{
			{
				ID:   "elm1",
				Type: "line",
				X:    25,
				Y:    30,
				Points: [][]float64{
					{0, 0}, {-10, 20},
				},
				StrokeWidth: 4,
			},
			{
				ID:   "elm2",
				Type: "line",
				X:    35,
				Y:    40,
				Points: [][]float64{
					{0, 0}, {-10, 20},
				},
				StrokeWidth: 1,
			},
			{
				ID:          "elm3",
				Type:        "ellipse",
				X:           35,
				Y:           55,
				Width:       20,
				Height:      30,
				Angle:       30,
				StrokeWidth: 3,
			},
		},
		AppState: inputdata.AppState{},
		Files:    nil,
		Tables:   nil,
		Paper: inputdata.Paper{
			Size:        "A4",
			Unit:        "mm",
			Orientation: "P",
		},
	}
	pdfData := NewPdfData(input)
	pdfData.GenPdf(nil, "test_gen_line.pdf")
}

func TestGenPdfFromFile(t *testing.T) {
	phMap := map[string]string{
		"user_name":      "occho 1",
		"tax_id":         "taxid1234",
		"invoice_suffix": "edv",
		"total":          "10,000.00",
		"total_price":    "1000000",
	}
	inputJson, err := inputdata.ReadData("../../resource/json/config.json")
	if err != nil {
		log.Fatal(err)
	}
	pdfDataJson := NewPdfData(inputJson)
	pdfDataJson.GenPdf(phMap, "test_gen_invoice_json.pdf")

	inputYaml, err := inputdata.ReadData("../../resource/yaml/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	pdfDataYaml := NewPdfData(inputYaml)
	pdfDataYaml.GenPdf(phMap, "test_gen_invoice_yaml.pdf")
}
